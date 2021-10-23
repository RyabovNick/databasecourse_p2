package route

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/RyabovNick/databasecourse_2/golang/godb/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/golang-jwt/jwt"
)

type DBInterface interface {
	CreateUser(db.User) (db.User, error)
	Auth(db.User) (db.User, error)
}

type Route struct {
	DB DBInterface
}

func NewRouter(ro Route) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Post("/signup", func(rw http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		res, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		var user db.User
		if err := json.Unmarshal(res, &user); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		user, err = ro.DB.CreateUser(user)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		token, err := GenToken(user)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.Write(token) //nolint
	})

	router.Post("/auth", func(rw http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		res, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		var user db.User
		if err := json.Unmarshal(res, &user); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		user, err = ro.DB.Auth(user)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		token, err := GenToken(user)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.Write(token) //nolint
	})

	return router
}

// GenToken generates token for user
func GenToken(user db.User) ([]byte, error) {
	// todo: add RSA256 keys
	token := jwt.NewWithClaims(jwt.SigningMethodNone, jwt.MapClaims{
		"login": user.Login,
	})
	t, err := token.SigningString()
	if err != nil {
		return nil, fmt.Errorf("sign token: %w", err)
	}

	type TokenResp struct {
		Token string `json:"token"`
	}

	tok := TokenResp{
		Token: t,
	}

	res, err := json.Marshal(tok)
	if err != nil {
		return nil, fmt.Errorf("marshal token: %w", err)
	}

	return res, nil
}
