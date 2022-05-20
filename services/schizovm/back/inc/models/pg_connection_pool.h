#pragma once

#include <memory>
#include <mutex>
#include <queue>
#include <condition_variable>

#include "models/pg_connection.h"

namespace svm {
    class pg_connection_pool;

    class pg_connection_guard {
    public:
        pg_connection_guard(std::shared_ptr<pg_connection> connection, pg_connection_pool *owner);

        pg_connection_guard(const pg_connection_guard &other) = delete;

        pg_connection_guard(pg_connection_guard &&other) = default;

        pg_connection_guard &operator=(const pg_connection_guard &other) = delete;

        pg_connection_guard &operator=(pg_connection_guard &&other) = delete;

        ~pg_connection_guard();

        const std::shared_ptr<pg_connection> connection;
    private:
        pg_connection_pool *owner;
    };

    class pg_connection_pool {
    public:
        pg_connection_pool(const pg_connection_config &config, int connectionCount);

        std::shared_ptr<pg_connection> get();

        void free(const std::shared_ptr<pg_connection> &connection);

        pg_connection_guard get_guarded();

    private:
        std::queue<std::shared_ptr<pg_connection>> connections;
        std::condition_variable free_connection_notifier;
        std::mutex mutex;
    };
}
