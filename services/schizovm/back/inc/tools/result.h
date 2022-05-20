#pragma once

#include <string>
#include <utility>
#include <functional>

namespace svm {
    template<class T>
    class result {
    public:
        const bool success;
        const T value;
        const std::string message;

        static result of_success(T value) {
            return {true, std::move(value), ""};
        }

        static result of_success(T value, std::string message) {
            return {true, std::move(value), std::move(message)};
        }

        static result of_error() {
            return {false, T(), ""};
        }

        static result of_error(std::string message) {
            return {false, T(), std::move(message)};
        }

        template <class TNext>
        result<TNext> then(std::function<result<TNext>(const T&)> map) {
            if (!success) {
                return result<TNext>::of_error(message);
            }

            return map(value);
        }

    private:
        result(bool success, T value, std::string message)
                : success(success), value(std::move(value)), message(std::move(message)) {
        }
    };
}
