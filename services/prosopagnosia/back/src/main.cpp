#include "crow/crow_all.h"

#include <iostream>
#include <string>
#include <random>
#include <sstream>

#include "services/fs.h"
#include "services/demo_service.h"
#include "options/options_parser.h"
#include "models/pg_connection_pool.h"
#include "tools/strings.h"

using namespace svm;
using namespace crow;

typedef App<CookieParser> App;

const char *NAME_HEADER = "X-Svm-Name";
const char *KEY_HEADER = "X-Svm-Key";
const char *AUTHOR_HEADER = "X-Svm-Author";
const char *SECRET_HEADER = "X-Svm-Secret";

response create_demo(const request &req, demo_service &demo_service) {
    auto name = req.get_header_value(NAME_HEADER);
    auto author = req.get_header_value(AUTHOR_HEADER);
    auto secret = req.get_header_value(SECRET_HEADER);

    if (name.empty() || secret.empty() || author.empty()
        || !is_base64_encoded_string(name) || !is_base64_encoded_string(author) || !is_base64_encoded_string(secret)) {
        return response(BAD_REQUEST);
    }

    if (req.get_header_value("Content-Type").find("multipart/form-data") == std::string::npos) {
        return response(BAD_REQUEST);
    }

    crow::multipart::message msg(req);

    if (msg.parts.size() != 1) {
        return response(BAD_REQUEST);
    }

    auto &part = msg.parts[0];
    std::string filename;

    std::vector<uint8_t> body;
    body.resize(part.body.size());
    std::copy_n(part.body.begin(), part.body.size(), body.begin());

    CROW_LOG_INFO << "rom size is " << body.size();

    auto demo = demo_service.create(name, author, secret, body);

    if (!demo.success) {
        return response(CONFLICT);
    }

    auto res = response(OK);

    res.set_header(KEY_HEADER, demo.value->key);

    return res;
}

response get_demo(const request &req, demo_service &demo_service) {
    auto name = req.get_header_value(NAME_HEADER);
    auto key = req.get_header_value(KEY_HEADER);

    if (name.empty() || key.empty() || !is_base64_encoded_string(name) || !is_base64_encoded_string(key)) {
        return response(BAD_REQUEST);
    }

    auto demo = demo_service.get(name, key);

    if (!demo.success) {
        return response(UNAUTHORIZED);
    }

    json::wvalue data(
            {
                    {"name",       demo.value->name},
                    {"author",     demo.value->author},
                    {"created_at", demo.value->created_at},
                    {"secret",     demo.value->secret},
                    {"rom_path",   demo.value->rel_rom_path().c_str()}
            });

    return {status::OK, data};
}

response list_demos(const request &req, demo_service &demo_service) {
    if (req.url_params.get("page_num") == nullptr || req.url_params.get("page_size") == nullptr) {
        return response(BAD_REQUEST);
    }

    auto page_num = std::stoi(std::string("0") + req.url_params.get("page_num"));
    auto page_size = std::stoi(std::string("0") + req.url_params.get("page_size"));
    auto list = demo_service.list(page_num, page_size);

    if (!list.success) {
        return response(INTERNAL_SERVER_ERROR);
    }

    std::vector<json::wvalue> jsoned;

    for (const auto &demo: list.value) {
        json::wvalue data(
                {
                        {"name",       demo->name},
                        {"author",     demo->author},
                        {"created_at", demo->created_at},
                        {"rom_path",   demo->rel_rom_path().c_str()}
                });

        jsoned.push_back(data);
    }

    json::wvalue data({{"demos", jsoned}});

    return {status::OK, data};
}

int main(int argc, char **argv) {
    try {
        auto args = command_line_parser::parse(argc, argv);

        ::App app;

        pg_connection_config pg_config = {
                get_env("POSTGRES_HOST"),
                get_int_env("POSTGRES_PORT"),
                get_env("POSTGRES_DB"),
                get_env("POSTGRES_USER"),
                get_env("POSTGRES_PASSWORD")
        };

        auto pg_pool = std::make_shared<pg_connection_pool>(pg_config, 10);
        auto storage_path = std::filesystem::path(get_env("SVM_STORAGE_PATH"));

        CROW_LOG_INFO << "storage will be in " << storage_path;

        auto fs = std::make_shared<svm::fs>(storage_path);
        auto demo_service = std::make_shared<svm::demo_service>(pg_pool, fs,
                                                                str_or_uuid4(get_env("SVM_SECRET", false)));

        CROW_ROUTE(app, "/api/demo").methods(HTTPMethod::Post, HTTPMethod::Get)(
                [&demo_service](const request &req) {
                    if (req.method == HTTPMethod::Post) {
                        return create_demo(req, *demo_service);
                    }

                    return get_demo(req, *demo_service);
                });

        CROW_ROUTE(app, "/api/demo/list").methods(HTTPMethod::Get)(
                [&demo_service](const request &req) {
                    return list_demos(req, *demo_service);
                });

        app.bindaddr(args.address).port(args.port).multithreaded().run();

    } catch (std::exception &e) {
        CROW_LOG_CRITICAL << "Unhandled exception: " << e.what();

        CROW_LOG_CRITICAL << "errno = " << errno << " (" << strerror(errno) << ")";

        return EXIT_FAILURE;
    } catch (...) {
        CROW_LOG_CRITICAL << "Unknown unhandled exception occurred";

        CROW_LOG_CRITICAL << "errno = " << errno << " (" << strerror(errno) << ")";

        return EXIT_FAILURE;
    }

    return 0;
}
