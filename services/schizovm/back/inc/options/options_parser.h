#pragma once

#include <string>

namespace svm {
    struct server_options {
    public:
        const std::string address;
        const int port;
        const std::string keys_secret;
    };

    class command_line_parser {
    public:
        static server_options parse(int argc, char **argv);
    };
}
