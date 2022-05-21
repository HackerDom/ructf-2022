package main

import (
	"snake/db"
	"snake/http"
)

func main() {
	db.Migrate()
	http.StartServ()
}
