package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

// TODO:
// 1. Создать отдельный пакет с БД, где будет подключение
// с создание connection pool. И подключения будут активные
// Например min: 1, max 3 подключений
// Также в этом пакете будут запросы к БД
// 2. Вынести routing в отдельный пакет
// 3. В main останется только создание объекта из пакет с БД
//    и создание routing, который прокидывается в
// http.ListenAndServe

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	connURI := "postgres://kalmykova:123123123@95.217.232.188:7777/kalmykova"

	router.Get("/menu/{id}", func(w http.ResponseWriter, r *http.Request) {
		conn, err := pgx.Connect(r.Context(), connURI)
		if err != nil {
			// TODO: hide error msg and log error
			// here and any other err check
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer conn.Close(r.Context())

		var name string
		if err := conn.QueryRow(r.Context(), `
		SELECT name 
		FROM menu
		WHERE id = $1
		`, chi.URLParam(r, "id")).Scan(&name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(name))
	})

	router.Get("/menu", func(w http.ResponseWriter, r *http.Request) {
		db, err := sqlx.Connect("pgx", connURI)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		var menu []Menu
		if err := db.SelectContext(r.Context(), &menu, `
		SELECT id, name, price
		FROM menu
		`); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res, err := json.Marshal(menu)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	})

	http.ListenAndServe(":3000", router)
}

type Menu struct {
	ID    int     `db:"id" json:"id"`
	Name  string  `db:"name" json:"name"`
	Price float64 `db:"price" json:"price"`
}
