package user

import (
	"encoding/json"
	"fmt"

	"github.com/golang-jwt/jwt"
)

type Rights string

const (
	Read       Rights = "r"
	Write      Rights = "w"
	AdminRead  Rights = "ar"
	AdminWrite Rights = "aw"
	Owner      Rights = "o"
)

type User struct {
	ID       string `db:"id" json:"id"`
	Login    string `db:"login" json:"login"`
	Password string `db:"password" json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

// GenerateToken generates token for user
func (u User) GenerateToken() ([]byte, error) {
	// todo: add RSA256 keys
	token := jwt.NewWithClaims(jwt.SigningMethodNone, jwt.MapClaims{
		"id":    u.ID,
		"login": u.Login,
	})
	t, err := token.SigningString()
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
