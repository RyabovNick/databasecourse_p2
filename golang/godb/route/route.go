package route

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/RyabovNick/databasecourse_2/golang/godb/db"
	"github.com/RyabovNick/databasecourse_2/golang/godb/errors"
	"github.com/RyabovNick/databasecourse_2/golang/godb/middleware/auth"
	"github.com/RyabovNick/databasecourse_2/golang/godb/models/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type DBInterface interface {
	CreateUser(context.Context, user.User) (user.User, error)
	Auth(context.Context, user.User) (user.User, error)
	CreateTodoList(context.Context, db.TodoList, user.User) (db.TodoList, error)
}

type Route struct {
	DB DBInterface
}

func NewRouter(ro Route) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(auth.Auth())

	router.Post("/signup", func(rw http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var user user.User
		if err := UnmarshalBody(r.Body, &user); err != nil {
			errors.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		user, err := ro.DB.CreateUser(r.Context(), user)
		if err != nil {
			errors.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		token, err := user.GenerateToken()
		if err != nil {
			errors.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.Write(token) //nolint
	})

	router.Post("/auth", func(rw http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var user user.User
		if err := UnmarshalBody(r.Body, &user); err != nil {
			errors.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		user, err := ro.DB.Auth(r.Context(), user)
		if err != nil {
			errors.Error(rw, err.Error(), http.StatusUnauthorized)
			return
		}

		token, err := user.GenerateToken()
		if err != nil {
			errors.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.Write(token) //nolint
	})

	router.Post("/create/todo_list", func(rw http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var todoList db.TodoList
		if err := UnmarshalBody(r.Body, &todoList); err != nil {
			errors.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		todoList, err := ro.DB.CreateTodoList(r.Context(), todoList, user.User{
			ID: r.Context().Value("id").(string),
		})
		if err != nil {
			errors.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "application/json")

		b, err := json.Marshal(todoList)
		if err != nil {
			errors.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		rw.Write(b)
	})

	return router
}

func UnmarshalBody(r io.Reader, v interface{}) error {
	res, err := ioutil.ReadAll(r)
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}

	if err := json.Unmarshal(res, v); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	return nil
}
