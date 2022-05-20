#pragma once

#include <string>
#include <memory>

#include <libpq-fe.h>

namespace svm {
    struct pg_connection_config {
        const std::string host;
        const int port;
        const std::string db_name;
        const std::string user;
        const std::string password;
    };

    class pg_connection {
    public:
        explicit pg_connection(const pg_connection_config &config);

        pg_connection(const pg_connection &other) = delete;

        pg_connection(pg_connection &&other) = delete;

        pg_connection &operator=(const pg_connection &other) = delete;

        pg_connection &operator=(pg_connection &&other) = delete;

        [[nodiscard]] std::shared_ptr<PGconn> get_connection() const;

    private:
        std::shared_ptr<PGconn> connection;
    };
}
