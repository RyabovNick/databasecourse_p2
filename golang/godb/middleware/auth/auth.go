package auth

import (
	"context"
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
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			token, err := jwt.ParseWithClaims(s[1], &user.WithClaims{}, func(t *jwt.Token) (interface{}, error) {
				return user.AuthToken, nil
			})
			if err != nil {
				log.Default().Println("token parse", err)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			if !token.Valid {
				log.Default().Println("not valid token")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			claims, ok := token.Claims.(*user.WithClaims)
			if !ok {
				log.Default().Println("claims", ok)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			r = r.WithContext(context.WithValue(r.Context(), user.CtxKey(), claims.ToUser()))
			next.ServeHTTP(w, r)
		})
	}
}
