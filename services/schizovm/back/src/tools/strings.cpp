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

static const char *base64_chars[2] = {
        "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
        "abcdefghijklmnopqrstuvwxyz"
        "0123456789"
        "+/",

        "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
        "abcdefghijklmnopqrstuvwxyz"
        "0123456789"
        "-_"};

static unsigned int pos_of_char(const unsigned char chr) {
    if (chr >= 'A' && chr <= 'Z') return chr - 'A';
    else if (chr >= 'a' && chr <= 'z') return chr - 'a' + ('Z' - 'A') + 1;
    else if (chr >= '0' && chr <= '9') return chr - '0' + ('Z' - 'A') + ('z' - 'a') + 2;
    else if (chr == '+' || chr == '-')
        return 62;
    else if (chr == '/' || chr == '_') return 63;
    else
        throw std::runtime_error("Input is not valid base64-encoded data.");
}

std::string base64_encode(unsigned char const *bytes_to_encode, size_t in_len, bool url) {

    size_t len_encoded = (in_len + 2) / 3 * 4;

    unsigned char trailing_char = url ? '.' : '=';

    const char *base64_chars_ = base64_chars[url];

    std::string ret;
    ret.reserve(len_encoded);

    unsigned int pos = 0;

    while (pos < in_len) {
        ret.push_back(base64_chars_[(bytes_to_encode[pos + 0] & 0xfc) >> 2]);

        if (pos + 1 < in_len) {
            ret.push_back(
                    base64_chars_[((bytes_to_encode[pos + 0] & 0x03) << 4) + ((bytes_to_encode[pos + 1] & 0xf0) >> 4)]);

            if (pos + 2 < in_len) {
                ret.push_back(base64_chars_[((bytes_to_encode[pos + 1] & 0x0f) << 2) +
                                            ((bytes_to_encode[pos + 2] & 0xc0) >> 6)]);
                ret.push_back(base64_chars_[bytes_to_encode[pos + 2] & 0x3f]);
            } else {
                ret.push_back(base64_chars_[(bytes_to_encode[pos + 1] & 0x0f) << 2]);
                ret.push_back(trailing_char);
            }
        } else {

            ret.push_back(base64_chars_[(bytes_to_encode[pos + 0] & 0x03) << 4]);
            ret.push_back(trailing_char);
            ret.push_back(trailing_char);
        }

        pos += 3;
    }


    return ret;
}


template<typename String, unsigned int line_length>
static std::string encode_with_line_breaks(String s) {
    return insert_linebreaks(base64_encode(s, false), line_length);
}

template<typename String>
static std::string encode(String s, bool url) {
    return base64_encode(reinterpret_cast<const unsigned char *>(s.data()), s.length(), url);
}

template<typename String>
static std::string decode(String encoded_string, bool remove_linebreaks) {
    if (encoded_string.empty()) return {};

    if (remove_linebreaks) {

        std::string copy(encoded_string);

        copy.erase(std::remove(copy.begin(), copy.end(), '\n'), copy.end());

        return base64_decode(copy, false);
    }

    size_t length_of_string = encoded_string.length();
    size_t pos = 0;

    size_t approx_length_of_decoded_string = length_of_string / 4 * 3;
    std::string ret;
    ret.reserve(approx_length_of_decoded_string);

    while (pos < length_of_string) {
        size_t pos_of_char_1 = pos_of_char(encoded_string[pos + 1]);

        ret.push_back(static_cast<std::string::value_type>(((pos_of_char(encoded_string[pos + 0])) << 2) +
                                                           ((pos_of_char_1 & 0x30) >> 4)));

        if ((pos + 2 < length_of_string) &&
            encoded_string[pos + 2] != '=' &&
            encoded_string[pos + 2] != '.'
                ) {
            unsigned int pos_of_char_2 = pos_of_char(encoded_string[pos + 2]);
            ret.push_back(static_cast<std::string::value_type>(((pos_of_char_1 & 0x0f) << 4) +
                                                               ((pos_of_char_2 & 0x3c) >> 2)));

            if ((pos + 3 < length_of_string) &&
                encoded_string[pos + 3] != '=' &&
                encoded_string[pos + 3] != '.') {
                ret.push_back(static_cast<std::string::value_type>(((pos_of_char_2 & 0x03) << 6) +
                                                                   pos_of_char(encoded_string[pos + 3])));
            }
        }

        pos += 4;
    }

    return ret;
}

std::string svm::base64_decode(std::string const &s, bool remove_linebreaks) {
    return decode(s, remove_linebreaks);
}

std::string svm::base64_encode(std::string const &s, bool url) {
    return encode(s, url);
}
