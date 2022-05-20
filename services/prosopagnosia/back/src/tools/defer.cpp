#include <cstdlib>
#include <iostream>

#include "tools/defer.h"

using namespace svm;

static thread_local std::stack<defer *> *this_thread_defers;

static std::mutex m;
static bool atExitInitialized = false;
static std::vector<std::stack<defer *> *> all_threads_defers_list;

void do_defer_all_on_exit() {
    std::scoped_lock<std::mutex> lock(m);

    for (auto &thread_defer: all_threads_defers_list) {
        while (!thread_defer->empty()) {
            thread_defer->top()->execute_all();
            thread_defer->pop();
        }
    }
}

std::stack<defer *> &get_thread_defers() {
    if (this_thread_defers == nullptr) {
        this_thread_defers = new std::stack<defer *>();

        {
            std::scoped_lock<std::mutex> lock(m);

            all_threads_defers_list.push_back(this_thread_defers);

            if (!atExitInitialized) {
                atExitInitialized = true;
                std::atexit(do_defer_all_on_exit);
            }
        }
    }

    return *this_thread_defers;
}

defer::defer() {
    get_thread_defers().push(this);
}

defer::~defer() {
    execute_all();

    get_thread_defers().pop();
}

void defer::execute_all() {
    for (auto &a: actions) {
        a();
    }

    actions.clear();
}

void defer::add(std::function<void()> &&action) {
    actions.push_back(std::forward<std::function<void()>>(action));
}
