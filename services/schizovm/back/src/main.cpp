#include "crow/crow_all.h"


int main() {
    crow::SimpleApp app;

    CROW_ROUTE(app, "/api/hello")([](){
        return "Hello world";
    });

    app.port(31337).run();
    
    return 0;
}
