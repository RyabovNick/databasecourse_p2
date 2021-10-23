// package db устанавливает соединение с БД
// и реализует методы - запросы к БД
package db

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrWrongPassword = fmt.Errorf("wrong password")
)

type DB struct {
	conn *sqlx.DB
}

func NewDB() (*DB, error) {
	connURI := "postgres://zhavoronok:123123123@95.217.232.188:7777/zhavoronok"

	db, err := sqlx.Connect("pgx", connURI)
	if err != nil {
		return nil, fmt.Errorf("sqlx connect: %w", err)
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(3)

	return &DB{
		conn: db,
	}, nil
}

type User struct {
	ID       string
	Login    string `db:"login" json:"login"`
	Password string `db:"password" json:"password"`
}

func (d *DB) CreateUser(user User) (User, error) {
	hashp, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, fmt.Errorf("hash: %w", err)
	}

	var id string
	if err := d.conn.Get(&id, `
	INSERT INTO users(login, password)
	VALUES ($1, $2)
	RETURNING id`, user.Login, string(hashp)); err != nil {
		return User{}, fmt.Errorf("create user: %w", err)
	}

	user.ID = id
	return user, nil
}

func (d *DB) Auth(user User) (User, error) {
	var hash string
	if err := d.conn.Get(&hash, `
	SELECT password
	FROM users
	WHERE login = $1`, user.Login); err != nil {
		return User{}, fmt.Errorf("auth: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(user.Password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return User{}, ErrWrongPassword
		}

		return User{}, fmt.Errorf("hash: %w", err)
	}

	return User{
		Login: user.Login,
	}, nil
}
