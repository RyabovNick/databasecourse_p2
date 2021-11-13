package user

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type Rights struct {
	ID   int
	Name string
}

var (
	Read       = Rights{1, "r"}
	Write      = Rights{2, "w"}
	AdminRead  = Rights{3, "ar"}
	AdminWrite = Rights{4, "aw"}
	Owner      = Rights{5, "o"}
	AuthToken  = []byte("secretKey")
)

func CheckRight(atleastHas Rights, has Rights) bool {
	// TODO

	return true
}

type User struct {
	ID       string `db:"id" json:"id"`
	Login    string `db:"login" json:"login"`
	Password string `db:"password" json:"password"`
}

type WithClaims struct {
	jwt.StandardClaims
	ID    string `json:"id"`
	Login string `json:"login"`
}

func (w *WithClaims) ToUser() User {
	return User{
		ID:    w.ID,
		Login: w.Login,
	}
}

type Token struct {
	Token string `json:"token"`
}

type Key struct {
	K string
}

func CtxKey() Key {
	return Key{K: "id"}
}

// GenerateToken generates token for user
func (u User) GenerateToken() ([]byte, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, WithClaims{
		ID:    u.ID,
		Login: u.Login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	})
	t, err := token.SignedString(AuthToken)
	if err != nil {
		return nil, fmt.Errorf("sign token: %w", err)
	}

	tok := Token{
		Token: t,
	}

	res, err := json.Marshal(tok)
	if err != nil {
		return nil, fmt.Errorf("marshal token: %w", err)
	}

	return res, nil
}

func FromCtx(ctx context.Context) User {
	return ctx.Value(CtxKey()).(User)
}
