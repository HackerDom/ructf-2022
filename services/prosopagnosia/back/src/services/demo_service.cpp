#include <utility>

#include "crow/crow_all.h"

#include "services/demo_service.h"
#include "tools/strings.h"
#include "tools/defer.h"
#include "tools/sha256.h"

using namespace svm;

demo_service::demo_service(std::shared_ptr<pg_connection_pool> pg_pool,
                           std::shared_ptr<svm::fs> fs,
                           std::string keys_secret)
        : pg_pool(std::move(pg_pool)), fs(std::move(fs)), keys_secret(std::move(keys_secret)) {
}

result<std::shared_ptr<demo>> demo_service::create(
        const std::string &name,
        const std::string &author,
        const std::string &secret,
        const std::vector<uint8_t> &rom) {
    defer defer;

    auto save_result = fs->save_next_file(rom);

    if (!save_result.success) {
        CROW_LOG_ERROR << "rom saving error: " << save_result.message;

        return result<std::shared_ptr<demo>>::of_error("cant save rom");
    }

    CROW_LOG_INFO << "rom saved at " << save_result.value;

    PGresult *result = nullptr;
    defer.add(([&result] {
        PQclear(result);
    }));

    auto guard = pg_pool->get_guarded();
    auto conn = guard.connection->get_connection().get();
    auto escaped_name = escape_psql(conn, name);
    auto escaped_author = escape_psql(conn, author);
    auto escaped_secret = escape_psql(conn, secret);
    auto escaped_key = escape_psql(conn, get_key(name));
    auto escaped_path = escape_psql(conn, save_result.value);

    auto query = format(
            hidden_str("insert into demos values (%s, %s, %s, %s, %s, default) returning *;"),
            escaped_name.c_str(),
            escaped_author.c_str(),
            escaped_secret.c_str(),
            escaped_key.c_str(),
            escaped_path.c_str()
    );

    result = PQexec(conn, query.c_str());

    if (PQresultStatus(result) != PGRES_TUPLES_OK || PQntuples(result) != 1) {
        auto msg = std::string(PQresultErrorMessage(result));

        CROW_LOG_ERROR << "creation of demo '" << name << "' failed with: " << msg;

        return svm::result<std::shared_ptr<demo>>::of_error("possible name conflict");
    }

    return svm::result<std::shared_ptr<demo>>::of_success(demo::from_pg_result(result));
}

result<std::shared_ptr<demo>> demo_service::get(const std::string &name, const std::string &key) {
    defer defer;

    PGresult *result = nullptr;
    defer.add(([&result] {
        PQclear(result);
    }));

    auto guard = pg_pool->get_guarded();
    auto conn = guard.connection->get_connection().get();

    auto escaped_name = escape_psql(conn, name);
    auto escaped_key = escape_psql(conn, key);

    auto query = format(
            hidden_str("select * from demos where name=%s and key=%s;"),
            escaped_name.c_str(),
            escaped_key.c_str()
    );

    result = PQexec(conn, query.c_str());

    if (PQresultStatus(result) != PGRES_TUPLES_OK || PQntuples(result) != 1) {
        return svm::result<std::shared_ptr<demo>>::of_error("invalid login or password");
    }

    return svm::result<std::shared_ptr<demo>>::of_success(demo::from_pg_result(result));
}

result<std::vector<std::shared_ptr<demo>>> demo_service::list(int page_num, int page_size) {
    svm::defer defer;
    PGresult *result = nullptr;
    defer.add(([&result] {
        PQclear(result);
    }));

    if (page_size <= 0) {
        page_size = 10;
    }

    page_size = std::min(page_size, 1000);

    if (page_num < 0) {
        page_num = 0;
    }

    auto guard = pg_pool->get_guarded();
    auto conn = guard.connection->get_connection().get();

    auto query = format(
            hidden_str("select * from demos order by created_at desc limit %d offset %d;"),
            page_size, page_num * page_size
    );

    result = PQexec(conn, query.c_str());

    if (PQresultStatus(result) != PGRES_TUPLES_OK) {
        auto err = std::string(PQresultErrorMessage(result));

        CROW_LOG_ERROR << "demos listing failed: " << err;

        return svm::result<std::vector<std::shared_ptr<demo>>>::of_error();
    }

    std::vector<std::shared_ptr<demo>> users;

    int n = PQntuples(result);

    for (int i = 0; i < n; ++i) {
        auto demo = demo::from_pg_result(result, i);

        users.push_back(demo);
    }

    return svm::result<std::vector<std::shared_ptr<demo>>>::of_success(users);
}


std::string demo_service::get_key(const std::string &name) {
    auto name_decoded = base64_decode(name);

    SHA256 sha;
    sha.update(name_decoded.value);
    sha.update(keys_secret);

    auto digest = sha.digest();
    std::string sha256 = SHA256::toString(digest.get());

    return base64_encode(sha256);
}
