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
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{
		DatabaseName: "zhavoronok",
		SchemaName:   "public",
	})
	if err != nil {
		return nil, fmt.Errorf("migrate instance: %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"zhavoronok", driver)
	if err != nil {
		return nil, fmt.Errorf("migrate: %w", err)
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, fmt.Errorf("migrate up: %w", err)
	}

	return &DB{
		conn: db,
	}, nil
}

func (d *DB) InTx(ctx context.Context, isolation sql.IsolationLevel, f func(tx sqlx.Tx) error) error {
	tx, err := d.conn.BeginTxx(ctx, &sql.TxOptions{
		Isolation: isolation,
	})
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	if err := f(*tx); err != nil {
		if errRoll := tx.Rollback(); errRoll != nil {
			return fmt.Errorf("rollback tx: %v (error: %w)", errRoll, err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
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
	// DeletedAt   time.Time `db:"deleted_at" json:"deleted_at"`
}

type UserRights struct {
	UsersID     string      `db:"users_id" json:"users_id"`
	TodoListsID string      `db:"todo_lists_id" json:"todo_lists_id"`
	Rights      user.Rights `db:"rights" json:"rights"`
}

// CreateTodoList creates todo list and sets rights as owner
func (d *DB) CreateTodoList(ctx context.Context, t TodoList, u user.User) (TodoList, error) {
	var todoList TodoList

	if err := d.InTx(ctx, sql.LevelReadCommitted, func(tx sqlx.Tx) error {
		if err := tx.Get(&todoList, `
		INSERT INTO todo_lists (title, created_by)
		VALUES ($1, $2)
		RETURNING id, title, created_at, created_by`, t.Title, u.ID); err != nil {
			return fmt.Errorf("insert todo_lists: %w", err)
		}

		if _, err := tx.Exec(`INSERT INTO user_rights 
		(users_id, todo_lists_id, rights)
		VALUES ($1, $2, $3)`, u.ID, todoList.ID, user.Owner); err != nil {
			return fmt.Errorf("insert user_rights: %w", err)
		}

		return nil
	}); err != nil {
		return TodoList{}, err
	}

	return todoList, nil
}

// GetRights returns user right for todo_list
func (d *DB) GetRights(ctx context.Context, todoListID, userID string) (user.Rights, error) {
	var right int
	if err := d.conn.GetContext(ctx, &right, `
	SELECT rights
	FROM user_rights
	WHERE todo_lists_id = $1 AND users_id = $2 
	`, todoListID, userID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user.NoRights, nil
		}

		return user.NoRights, fmt.Errorf("get user rights: %w", err)
	}

	return user.IntToRights(right), nil
}

// AvailableTodoLists returns available todo lists for user
func (d *DB) AvailableTodoLists(ctx context.Context, userID string) ([]TodoList, error) {
	var tl []TodoList
	if err := d.conn.SelectContext(ctx, &tl, `
	SELECT tl.id, tl.title, tl.created_at, tl.created_by
	FROM todo_lists tl
	INNER JOIN user_rights ur ON  tl.created_by = ur.users_id 
		AND tl.id = ur.todo_lists_id
	WHERE tl.created_by = $1
	`, userID); err != nil {
		return nil, fmt.Errorf("get todo lists: %w", err)
	}

	return tl, nil
}

func (d *DB) GetTodoListTodo(ctx context.Context, todoListID string) ([]Todo, error) {
	var td []Todo
	if err := d.conn.SelectContext(ctx, &td, `
	SELECT id, title, description, checked, 
		todo_lists_id, created_at, created_by, 
		updated_at, updated_by
	FROM todos
	WHERE todo_lists_id = $1
	`, todoListID); err != nil {
		return nil, fmt.Errorf("get todo: %w", err)
	}

	return td, nil
}

func (d *DB) CreateRights(ctx context.Context, rights UserRights) error {
	if _, err := d.conn.ExecContext(ctx, `
	INSERT INTO user_rights (users_id, todo_lists_id, rights)
	VALUES ($1, $2, $3)`,
		rights.UsersID, rights.TodoListsID, rights.Rights); err != nil {
		return fmt.Errorf("create user rights: %w", err)
	}

	return nil
}
