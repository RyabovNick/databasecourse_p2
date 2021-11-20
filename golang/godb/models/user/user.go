package user

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type Rights int

var (
	NoRights   Rights = 0
	Read       Rights = 1 // Только чтение
	Write      Rights = 2 // Чтение и запись
	AdminRead  Rights = 3 // Чтение, запись и возможность дать права на чтение
	AdminWrite Rights = 4 // Чтение, запись и возможность дать права на чтение и запись
	Owner      Rights = 5 // Любые действия
	AuthToken         = []byte("secretKey")
)

func IntToRights(v int) Rights {
	switch v {
	case 1:
		return Read
	case 2:
		return Write
	case 3:
		return AdminRead
	case 4:
		return AdminWrite
	case 5:
		return Owner
	}

	return NoRights
}

func CheckRight(atleastHas Rights, has Rights) bool {
	return has >= atleastHas
}

func CanGiveRights(with Rights, give Rights) bool {
	if with == Owner {
		return true
	}

	if with == AdminWrite && (give == Read || give == Write) {
		return true
	}

	if with == AdminRead && give == Read {
		return true
	}

	return false
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
