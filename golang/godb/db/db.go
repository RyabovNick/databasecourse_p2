// package db устанавливает соединение с БД
// и реализует методы - запросы к БД
package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type DB struct {
	conn *sqlx.DB
}

func NewDB() (*DB, error) {
	connURI := "postgres://kalmykova:123123123@95.217.232.188:7777/kalmykova"

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

type Menu struct {
	ID    int     `db:"id" json:"id"`
	Name  string  `db:"name" json:"name"`
	Price float64 `db:"price" json:"price"`
}

func (d *DB) SelectMenu() ([]Menu, error) {
	var menu []Menu
	if err := d.conn.Select(&menu, `
	SELECT id, name, price
	FROM menu
	`); err != nil {
		return nil, fmt.Errorf("select menu: %w", err)
	}

	return menu, nil
}

func (d *DB) GetMenuByID(id string) (Menu, error) {
	var menu Menu
	if err := d.conn.Get(&menu, `
	SELECT id, name, price
	FROM menu
	WHERE id = $1`, id); err != nil {
		return Menu{}, fmt.Errorf("get menu: %w", err)
	}

	return menu, nil
}
