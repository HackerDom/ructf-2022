#pragma once

#include <sstream>
#include <string>
#include <vector>
#include <cstdlib>
#include <memory>
#include <stdexcept>

#include <libpq-fe.h>

#include "tools/result.h"

namespace svm {
    std::string get_env(const char *name, bool req = true);

    int get_int_env(const char *name);

    std::vector<std::string> split(const std::string &s, const std::string &delimiter);

    std::string &ltrim(std::string &s);

    std::string &rtrim(std::string &s);

    std::string &trim(std::string &s);

    std::string escape_psql(PGconn *conn, const std::string &s);

    std::string escape_psql(PGconn *conn, const char *s);

    std::string generate_uuid_v4();

    std::string str_or_uuid4(std::string s);

    std::string base64_encode(const std::string &s);

    result<std::string> base64_decode(const std::string &s);

    bool is_base64_encoded_string(const std::string &s);

    template<typename ... Args>
    std::string format(const char *format, Args ... args) {
        int sizeS = std::snprintf(nullptr, 0, format, args ...) + 1;
        if (sizeS <= 0) {
            throw std::runtime_error("Error during formatting.");
        }

        auto size = static_cast<size_t>(sizeS);
        auto buf = std::make_unique<char[]>(size);

        std::snprintf(buf.get(), size, format, args ...);

        return {buf.get(), buf.get() + size - 1};
    }

    template<size_t N>
    struct XorString {
    private:
        const char _key;
        std::array<char, N + 1> _encrypted;

        [[nodiscard]] constexpr char enc(char c) const {
            return c ^ _key;
        }

        [[nodiscard]] char dec(char c) const {
            return c ^ _key;
        }

    public:
        template<size_t... Is>
        constexpr XorString(const char *str, std::index_sequence<Is...>) : _key(0x7d), _encrypted{enc(str[Is])...} {
        }

        auto decrypt() {
            for (size_t i = 0; i < N; ++i) {
                _encrypted[i] = dec(_encrypted[i]);
            }
            _encrypted[N] = '\0';
            return _encrypted.data();
        }

#define hidden_str(s) []{ constexpr XorString<sizeof(s)/sizeof(char) - 1> expr( s, std::make_index_sequence< sizeof(s)/sizeof(char) - 1>() ); return expr; }().decrypt()
    };
}
