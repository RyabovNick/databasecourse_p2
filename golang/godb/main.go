package main

import (
	"net/http"

	"github.com/RyabovNick/databasecourse_2/golang/godb/db"
	"github.com/RyabovNick/databasecourse_2/golang/godb/route"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	conn, err := db.NewDB()
	if err != nil {
		panic(err)
	}

	router := route.NewRouter(route.Route{DB: conn})
	http.ListenAndServe(":3000", router)
}
