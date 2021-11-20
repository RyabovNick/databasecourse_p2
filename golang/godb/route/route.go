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
	GetRights(context.Context, string, string) (user.Rights, error)
	AvailableTodoLists(context.Context, string) ([]db.TodoList, error)
	GetTodoListTodo(context.Context, string) ([]db.Todo, error)
	CreateRights(context.Context, db.UserRights) error
}

type Route struct {
	DB DBInterface
}

func NewRouter(ro Route) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

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

	router.Group(func(rout chi.Router) {
		rout.Use(auth.Auth())

		rout.Post("/todo_list", func(rw http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			var todoList db.TodoList
			if err := UnmarshalBody(r.Body, &todoList); err != nil {
				errors.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}

			us := user.FromCtx(r.Context())

			todoList, err := ro.DB.CreateTodoList(r.Context(), todoList, user.User{
				ID:    us.ID,
				Login: us.Login,
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

		rout.Get("/todo_list", func(rw http.ResponseWriter, r *http.Request) {
			us := user.FromCtx(r.Context())

			tl, err := ro.DB.AvailableTodoLists(r.Context(), us.ID)
			if err != nil {
				errors.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}

			res, err := json.Marshal(tl)
			if err != nil {
				errors.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}

			rw.Header().Set("Content-Type", "application/json")
			rw.Write(res)
		})

		rout.Get("/todo_list/{id}", func(rw http.ResponseWriter, r *http.Request) {
			us := user.FromCtx(r.Context())
			tlID := chi.URLParam(r, "id")

			// Check rights
			has, err := ro.DB.GetRights(r.Context(), tlID, us.ID)
			if err != nil {
				errors.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}

			if ok := user.CheckRight(user.Read, has); !ok {
				errors.Error(rw, "Access Forbidden", http.StatusForbidden)
				return
			}

			// Get TodoList
			tl, err := ro.DB.GetTodoListTodo(r.Context(), tlID)
			if err != nil {
				errors.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}

			res, err := json.Marshal(tl)
			if err != nil {
				errors.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}

			rw.Header().Set("Content-Type", "application/json")
			rw.Write(res)
		})

		rout.Post("/user_rights", func(rw http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()

			var userRights db.UserRights
			if err := UnmarshalBody(r.Body, &userRights); err != nil {
				errors.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}

			us := user.FromCtx(r.Context())

			rights, err := ro.DB.GetRights(r.Context(), userRights.TodoListsID, us.ID)
			if err != nil {
				errors.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}

			if ok := user.CanGiveRights(rights, userRights.Rights); !ok {
				errors.Error(rw, "No rights", http.StatusForbidden)
				return
			}

			if err := ro.DB.CreateRights(r.Context(), userRights); err != nil {
				errors.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}

			rw.WriteHeader(http.StatusOK)
		})
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
