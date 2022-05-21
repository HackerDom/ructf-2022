#include <iostream>
#include <cstdlib>

#include <boost/program_options.hpp>
#include <random>

#include "options/options_parser.h"

using namespace svm;

namespace po = boost::program_options;

server_options parse(int argc, char **argv) {
    po::options_description general_options("General options");
    std::string type, address, keys_secret;
    int port;
    general_options.add_options()
            ("help", "Show help")
            ("address", po::value<std::string>(&address)->required(), "Address to bind for listening")
            ("port", po::value<int>(&port)->required(), "Port to bind for listening");

    po::variables_map var_map;
    auto parsed = po::command_line_parser(argc, argv).options(general_options).allow_unregistered().run();
    po::store(parsed, var_map);

    if (var_map.count("help")) {
        std::cout << general_options << std::endl;

        std::exit(EXIT_SUCCESS);
    }

    po::notify(var_map);

    return {
            address,
            port
    };
}

server_options command_line_parser::parse(int argc, char **argv) {
    try {
        return ::parse(argc, argv);
    } catch (std::exception &e) {
        std::cout << e.what() << std::endl
                  << "See --help for help" << std::endl;
        std::exit(EXIT_FAILURE);
    } catch (...) {
        std::cout << "Unknown error" << std::endl
                  << "See --help for help" << std::endl;
        std::exit(EXIT_FAILURE);
    }
}
