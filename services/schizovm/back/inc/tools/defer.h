#pragma once

#include <mutex>
#include <vector>
#include <stack>
#include <functional>
#include <atomic>

namespace svm {
    class defer {
    public:
        defer();

        ~defer();

        defer(const defer &other) = delete;

        defer(defer &&other) = delete;

        defer &operator=(const defer &other) = delete;

        defer &operator=(defer &&other) = delete;

        void add(std::function<void()> &&action);

        template<class F, class ...Args>
        void operator()(F &&action, Args &&...args) {
            auto bind = std::bind(std::forward<F>(action), std::forward<Args>(args)...);

            Add([bind] { bind(); });
        }

        void execute_all();

        [[noreturn]] void *operator new(size_t) {
            throw std::runtime_error("defer mustn't be allocated on heap");
        }

        [[noreturn]] void operator delete(void*) {
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wexceptions"
            throw std::runtime_error("defer mustn't be allocated on heap");
#pragma clang diagnostic pop
        }

    private:
        std::vector<std::function<void()>> actions;
    };
}
