package auth

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/RyabovNick/databasecourse_2/golang/godb/models/user"
	"github.com/golang-jwt/jwt"
)

// Auth implements a simple middleware handler for bearer token
func Auth() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bearer := r.Header.Get("Authorization")
			s := strings.Split(bearer, " ")

			log.Default().Println("test")

			if len(s) != 2 {
				log.Default().Println("split")
				next.ServeHTTP(w, r) // TODO: not use middleware for auth & signup
				return
			}

			// TODO: decode and validate key
			b, err := jwt.Parse(s[1], func(t *jwt.Token) (interface{}, error) {
				// TODO: https://pkg.go.dev/github.com/golang-jwt/jwt#example-Parse-Hmac
				return nil, nil
			})
			if err != nil {
				// TODO
				log.Default().Println("decode", err)
				return
			}

			var user user.User
			if err := json.Unmarshal([]byte(b.Raw), &user); err != nil {
				log.Default().Println("unmarshal", err)
				return
			}

			r.WithContext(context.WithValue(r.Context(), "id", user.ID))
			next.ServeHTTP(w, r)
		})
	}
}
