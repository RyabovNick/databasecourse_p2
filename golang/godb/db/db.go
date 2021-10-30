// package db устанавливает соединение с БД
// и реализует методы - запросы к БД
package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/RyabovNick/databasecourse_2/golang/godb/models/user"
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

func (d *DB) CreateUser(ctx context.Context, u user.User) (user.User, error) {
	hashp, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return user.User{}, fmt.Errorf("hash: %w", err)
	}

	var id string
	if err := d.conn.GetContext(ctx, &id, `
	INSERT INTO users(login, password)
	VALUES ($1, $2)
	RETURNING id`, u.Login, string(hashp)); err != nil {
		return user.User{}, fmt.Errorf("create user: %w", err)
	}

	u.ID = id
	return u, nil
}

func (d *DB) Auth(ctx context.Context, got user.User) (user.User, error) {
	var us user.User
	if err := d.conn.GetContext(ctx, &us, `
	SELECT id, password, login
	FROM users
	WHERE login = $1`, got.Login); err != nil {
		return user.User{}, fmt.Errorf("auth: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(us.Password), []byte(got.Password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return user.User{}, ErrWrongPassword
		}

		return user.User{}, fmt.Errorf("hash: %w", err)
	}

	return user.User{
		ID:    us.ID,
		Login: us.Login,
	}, nil
}

type TodoList struct {
	ID        string    `db:"id" json:"id"`
	Title     string    `db:"title" json:"title"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	CreatedBy string    `db:"created_by" json:"created_by"`
}

type Todo struct {
	ID          string    `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	Checked     bool      `db:"checked" json:"checked"`
	TodoListsID string    `db:"todo_lists_id" json:"todo_lists_id"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	CreatedBy   string    `db:"created_by" json:"created_by"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy   string    `db:"updated_by" json:"updated_by"`
	DeletedAt   time.Time `db:"deleted_at" json:"deleted_at"`
}

type UserRights struct {
	UsersID     string `db:"users_id" json:"users_id"`
	TodoListsID string `db:"todo_lists_id" json:"todo_lists_id"`
	Rights      string `db:"rights" json:"rights"`
}

// CreateTodoList creates todo list and sets rights as owner
func (d *DB) CreateTodoList(ctx context.Context, t TodoList, u user.User) (TodoList, error) {
	tx, err := d.conn.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
	})
	if err != nil {
		return TodoList{}, fmt.Errorf("tx: %w", err)
	}

	var todoList TodoList
	if err := tx.Get(&todoList, `
	INSERT INTO todo_lists (title, created_by)
	VALUES ($1, $2)
	RETURNING id, title, created_at, created_by`, t.Title, u.ID); err != nil {
		if err := tx.Rollback(); err != nil {
			return TodoList{}, fmt.Errorf("rollback: %w", err)
		}
		return TodoList{}, fmt.Errorf("insert todo_lists: %w", err)
	}

	if _, err := tx.Exec(`INSERT INTO user_rights 
	(users_id, todo_lists_id, rights)
	VALUES ($1, $2, $3)`, u.ID, todoList.ID, user.Owner); err != nil {
		if err := tx.Rollback(); err != nil {
			return TodoList{}, fmt.Errorf("rollback: %w", err)
		}
		return TodoList{}, fmt.Errorf("insert user_rights: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return TodoList{}, fmt.Errorf("commit: %w", err)
	}

	return todoList, nil
}
