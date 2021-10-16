package route

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/RyabovNick/databasecourse_2/golang/godb/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type DBInterface interface {
	SelectMenu() ([]db.Menu, error)
	GetMenuByID(string) (db.Menu, error)
}

type Route struct {
	DB DBInterface
}

func NewRouter(ro Route) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/menu/{id}", func(w http.ResponseWriter, r *http.Request) {
		menu, err := ro.DB.GetMenuByID(chi.URLParam(r, "id"))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				http.Error(w, "not found", http.StatusNotFound)
				return
			}

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(menu.Name))
	})

	router.Get("/menu", func(w http.ResponseWriter, r *http.Request) {
		menu, err := ro.DB.SelectMenu()
		if err != nil {
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

	return router
}
