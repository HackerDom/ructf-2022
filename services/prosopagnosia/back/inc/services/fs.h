#pragma once

#include <vector>
#include <filesystem>

#include "tools/result.h"

namespace svm {
    class fs {
    public:
        explicit fs(std::filesystem::path basedir);

        std::filesystem::path get_path_for_next();

        result<std::filesystem::path> save_next_file(const std::vector<uint8_t> &data);

        const std::filesystem::path basedir;
    };
}
