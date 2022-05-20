#pragma once

#include <vector>

#include "tools/result.h"
#include "models/demo.h"
#include "services/fs.h"
#include "models/pg_connection_pool.h"

namespace svm {
    class demo_service {
    public:
        demo_service(std::shared_ptr<pg_connection_pool> pg_pool, std::shared_ptr<fs> fs, std::string keys_secret);

        result<std::shared_ptr<demo>> create(
                const std::string &name,
                const std::string &author,
                const std::string &secret,
                const std::vector<uint8_t> &rom
        );

        result<std::shared_ptr<demo>> get(const std::string &name, const std::string &key);

        result<std::vector<std::shared_ptr<demo>>> list(int page_num, int page_size);

    private:
        std::shared_ptr<pg_connection_pool> pg_pool;
        std::shared_ptr<fs> fs;
        std::string keys_secret;

        std::string get_key(const std::string &name);
    };
}