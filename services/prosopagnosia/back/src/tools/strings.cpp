#include <functional>
#include <cctype>
#include <locale>
#include <random>
#include <stdexcept>

#include "tools/strings.h"

using namespace svm;

std::string svm::get_env(const char *name, bool req) {
    auto *value = std::getenv(name);

    if (value == nullptr) {
        if (req) {
            throw std::runtime_error(format("environment variable '%s' required", name));
        }

        return {};
    }

    return {value};
}

int svm::get_int_env(const char *name) {
    return std::stoi(get_env(name));
}

std::vector<std::string> svm::split(const std::string &s, const std::string &delimiter) {
    size_t startIdx = 0, endIdx, delimSize = delimiter.length();
    std::string token;
    std::vector<std::string> res;

    while ((endIdx = s.find(delimiter, startIdx)) != std::string::npos) {
        token = s.substr(startIdx, endIdx - startIdx);
        startIdx = endIdx + delimSize;
        res.push_back(token);
    }

    res.push_back(s.substr(startIdx));

    return res;
}

std::string &svm::ltrim(std::string &s) {
    s.erase(s.begin(), std::find_if(s.begin(), s.end(),
                                    std::not1(std::ptr_fun<int, int>(std::isspace))));
    return s;
}

std::string &svm::rtrim(std::string &s) {
    s.erase(std::find_if(s.rbegin(), s.rend(),
                         std::not1(std::ptr_fun<int, int>(std::isspace))).base(), s.end());
    return s;
}

std::string &svm::trim(std::string &s) {
    return ltrim(rtrim(s));
}

std::string svm::escape_psql(PGconn *conn, const std::string &s) {
    char *escapedPtr = PQescapeLiteral(conn, s.c_str(), s.size());

    std::string copy(escapedPtr);

    PQfreemem(escapedPtr);

    return copy;
}

std::string svm::generate_uuid_v4() {
    static std::random_device rd;
    static std::mt19937 gen(rd());
    static std::uniform_int_distribution<> dis(0, 15);
    static std::uniform_int_distribution<> dis2(8, 11);

    std::stringstream ss;

    ss << std::hex;
    for (int i = 0; i < 8; i++) {
        ss << dis(gen);
    }
    ss << "-";
    for (int i = 0; i < 4; i++) {
        ss << dis(gen);
    }
    ss << "-4";
    for (int i = 0; i < 3; i++) {
        ss << dis(gen);
    }
    ss << "-";
    ss << dis2(gen);
    for (int i = 0; i < 3; i++) {
        ss << dis(gen);
    }
    ss << "-";
    for (int i = 0; i < 12; i++) {
        ss << dis(gen);
    }

    return ss.str();
}

std::string svm::str_or_uuid4(std::string s) {
    if (!s.empty()) {
        return s;
    }

    return generate_uuid_v4();
}

static const char *base64_alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";
static thread_local char *triples_map = nullptr;

static void init_triples_map() {
    if (triples_map != nullptr) {
        return;
    }

    // no need to deallocate - its static
    triples_map = new char[0xffffff * 4];

    for (int i = 0; i < 0xffffff; ++i) {
        auto base = i * 4;

        triples_map[base + 0] = base64_alpha[(i & 0xfc0000) >> 18];
        triples_map[base + 1] = base64_alpha[(i & 0x03f000) >> 12];
        triples_map[base + 2] = base64_alpha[(i & 0x000fc0) >> 6];
        triples_map[base + 3] = base64_alpha[(i & 0x00003f)];
    }
}

static const char *get_triplet_chars(uint32_t triplet) {
    init_triples_map();

    return triples_map + triplet * 4;
}

static int pos_of_char(const unsigned char chr) {
    if (chr >= 'A' && chr <= 'Z') {
        return chr - 'A';
    } else if (chr >= 'a' && chr <= 'z') {
        return chr - 'a' + ('Z' - 'A') + 1;
    } else if (chr >= '0' && chr <= '9') {
        return chr - '0' + ('Z' - 'A') + ('z' - 'a') + 2;
    } else if (chr == '+') {
        return 62;
    } else if (chr == '/') {
        return 63;
    }

    return -1;
}

std::string svm::base64_encode(const std::string &val) {
    std::string result;

    result.reserve((val.size() + 2) / 3 * 4 + 10);

    auto len = val.size();
    auto wl = len - len % 3;
    const char *end = &val[0] + wl;
    const char *curr = &val[0];

    while (curr < end) {
        auto triple = *(reinterpret_cast<const uint32_t *>(curr));
        triple = ((triple & 0xff) << 16) | (triple & 0xff00) | ((triple & 0xff0000) >> 16);

        auto chars = get_triplet_chars(triple);

        result.push_back(chars[0]);
        result.push_back(chars[1]);
        result.push_back(chars[2]);
        result.push_back(chars[3]);

        curr += 3;
    }

    if (wl + 1 == len) {
        auto last = val[len - 1];
        result.push_back(base64_alpha[(last & 0xfc) >> 2]);
        result.push_back(base64_alpha[(last & 0x3) << 4]);
        result.push_back('=');
        result.push_back('=');
    }

    if (wl + 2 == len) {
        auto last = val[len - 1];
        auto penultimate = val[len - 2];
        result.push_back(base64_alpha[(penultimate & 0xfc) >> 2]);
        result.push_back(base64_alpha[((penultimate & 0x3) << 4) | ((last & 0xf0) >> 4)]);
        result.push_back(base64_alpha[((last & 0xf) << 2)]);
        result.push_back('=');
    }

    return result;
}

result<std::string> svm::base64_decode(const std::string &val) {
    if (val.size() % 4 != 0 || std::count(val.begin(), val.end(), '=') > 2) {
        return result<std::string>::of_error("input is not valid base64-encoded data");
    }

    auto len = val.size();

    std::string decoded;
    decoded.reserve(len / 4 * 3);

    std::size_t i = 0;

    while (i < len) {
        auto c0 = val[i + 0];
        auto c1 = val[i + 1];
        auto c2 = val[i + 2];
        auto c3 = val[i + 3];

        auto p0 = pos_of_char(c0);
        auto p1 = pos_of_char(c1);

        decoded.push_back(static_cast<char>((p0 << 2) | (p1 & 0x30) >> 4));

        if (c2 != '=' && c3 != '=') {
            auto p2 = pos_of_char(c2);
            auto p3 = pos_of_char(c3);

            if (p0 == -1 || p1 == -1 || p2 == -1 || p3 == -1) {
                return result<std::string>::of_error("input is not valid base64-encoded data");
            }

            decoded.push_back(static_cast<char>(((p1 & 0x0f) << 4) | ((p2 & 0x3c) >> 2)));
            decoded.push_back(static_cast<char>((p2 & 0x03) << 6 | p3));
        } else {
            if (c2 == '=' && c3 != '=') {
                return result<std::string>::of_error("input is not valid base64-encoded data");
            }

            if (c2 != '=') {
                auto p2 = pos_of_char(c2);

                decoded.push_back(static_cast<char>(((p1 & 0x0f) << 4) | ((p2 & 0x3c) >> 2)));
            }
        }

        i += 4;
    }

    return result<std::string>::of_success(decoded);
}

bool svm::is_base64_encoded_string(const std::string &s) {
    return base64_decode(s).success;
}