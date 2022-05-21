#include "models/pg_connection_pool.h"

#include <utility>

using namespace svm;

pg_connection_guard::pg_connection_guard(std::shared_ptr<pg_connection> connection, pg_connection_pool *owner)
        : connection(std::move(connection)),
          owner(owner) {
}

pg_connection_guard::~pg_connection_guard() {
    owner->free(connection);
}

pg_connection_pool::pg_connection_pool(const pg_connection_config &config, int connectionCount) {
    std::scoped_lock<std::mutex> lock(mutex);

    connectionCount = std::max(1, connectionCount);

    for (auto i = 0; i < connectionCount; ++i) {
        connections.emplace(std::make_shared<pg_connection>(config));
    }
}

std::shared_ptr<pg_connection> pg_connection_pool::get() {
    std::unique_lock<std::mutex> lock(mutex);

    free_connection_notifier.wait(lock, [this] { return !connections.empty(); });

    auto connection = connections.front();
    connections.pop();

    return connection;
}

void pg_connection_pool::free(const std::shared_ptr<pg_connection> &connection) {
    std::unique_lock<std::mutex> lock(mutex);

    connections.push(connection);

    lock.unlock();

    free_connection_notifier.notify_one();
}

pg_connection_guard pg_connection_pool::get_guarded() {
    return {get(), this};
}
