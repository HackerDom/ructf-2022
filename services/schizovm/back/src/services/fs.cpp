#include <random>
#include <utility>
#include <fstream>
#include <cstdlib>
#include <cstring>

#include "services/fs.h"
#include "tools/strings.h"

using namespace svm;

static std::string uuid4() {
    std::random_device rd;
    std::mt19937 gen(rd());
    std::uniform_int_distribution<> dis(0, 15);
    std::uniform_int_distribution<> dis2(8, 11);
    std::stringstream ss;

    ss << std::hex;

    for (auto i = 0; i < 15; i++) {
        ss << dis(gen);
    }

    ss << dis2(gen);

    for (auto i = 0; i < 15; i++) {
        ss << dis(gen);
    }

    return ss.str();
}

fs::fs(std::filesystem::path basedir) : basedir(std::move(basedir)) {
    std::filesystem::create_directories(this->basedir);
}

std::filesystem::path fs::get_path_for_next() {
    auto name = uuid4();
    auto dir = basedir / name.substr(0, 2);

    std::filesystem::create_directories(dir);

    return dir / name.substr(2);
}

result<std::filesystem::path> fs::save_next_file(const std::vector<uint8_t> &data) {
    auto path = get_path_for_next();

    std::ofstream outfile(path, std::ios::out | std::ios::binary);

    if (!outfile.is_open()) {
        return result<std::filesystem::path>::of_error(
                format("opening '%s' for saving failed: %s", path.c_str(), strerror(errno))
        );
    }

    outfile.write(reinterpret_cast<const char *>(&data[0]), static_cast<int>(data.size()));

    outfile.close();

    return result<std::filesystem::path>::of_success(path);
}
