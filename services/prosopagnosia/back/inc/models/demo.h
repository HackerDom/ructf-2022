#pragma once

#include <memory>
#include <string>
#include <filesystem>

#include <libpq-fe.h>

namespace svm {
    class demo {
    public:
        demo(std::string name,
             std::string author,
             std::string secret,
             std::string key,
             std::filesystem::path rom_path,
             std::string created_at) :
                name(std::move(name)),
                author(std::move(author)),
                secret(std::move(secret)),
                key(std::move(key)),
                rom_path(std::move(rom_path)),
                created_at(std::move(created_at)) {
        }

        const std::string name;
        const std::string author;
        const std::string secret;
        const std::string key;
        const std::filesystem::path rom_path;
        const std::string created_at;

        static std::shared_ptr<demo> from_pg_result(PGresult *result, int row = 0) {
            auto name_col = PQfnumber(result, "name");
            auto author_col = PQfnumber(result, "author");
            auto secret_col = PQfnumber(result, "secret");
            auto key_col = PQfnumber(result, "key");
            auto rom_path_col = PQfnumber(result, "rom_path");
            auto created_at_col = PQfnumber(result, "created_at");

            if (name_col == -1
                || author_col == -1
                || secret_col == -1
                || key_col == -1
                || rom_path_col == -1
                || created_at_col == -1) {
                throw std::runtime_error("invalid pg result for demo");
            }

            return std::make_shared<demo>(
                    std::string(PQgetvalue(result, row, name_col)),
                    std::string(PQgetvalue(result, row, author_col)),
                    std::string(PQgetvalue(result, row, secret_col)),
                    std::string(PQgetvalue(result, row, key_col)),
                    std::filesystem::path(PQgetvalue(result, row, rom_path_col)),
                    std::string(PQgetvalue(result, row, created_at_col))
            );
        }

        std::filesystem::path rel_rom_path() {
            return "roms" / rom_path.parent_path().filename() / rom_path.filename();
        }
    };
}
