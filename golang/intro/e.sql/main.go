package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func main() {
	ctxwt, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	connUri := "postgres://kalmykova@95.217.232.188:7777/kalmykova"
	conn, err := pgx.Connect(ctxwt, connUri)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var name string                                                                  // ""
	err = conn.QueryRow(context.Background(), "select name from client").Scan(&name) // "..."
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			fmt.Fprintf(os.Stderr, "No found!")
			return
		}

		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return
	}

	fmt.Println(name)

	rows, err := conn.Query(context.Background(), `
	SELECT name, address
	FROM client
	`)
	if err != nil {
		// return fmt.Errorf("client query failed: %w", err)
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			name    string
			address string
		)

		if err := rows.Scan(&name, &address); err != nil {
			fmt.Fprintf(os.Stderr, "scan failed: %v\n", err)
			return
		}

		fmt.Println(name, address)
	}

	if rows.Err() != nil {
		fmt.Fprintf(os.Stderr, "scan failed: %v\n", err)
		return
	}

	db, errS := sqlx.Connect("pgx", connUri)
	if errS != nil {
		fmt.Fprintf(os.Stderr, "sqlx conn failed: %v\n", errS)
		return
	}

	var name1 string
	if err := db.Get(&name1, "SELECT name FROM client WHERE id = 1"); err != nil {
		fmt.Fprintf(os.Stderr, "sqlx get failed: %v\n", err)
		return
	}

	fmt.Println(name1)

	type Client struct {
		Id      int    `db:"id"`
		Name    string `db:"name"`
		Address string `db:"address"`
		Phone   string `db:"phone"`
	}

	cl := make([]Client, 0)
	if err := db.Select(&cl, "SELECT * FROM client"); err != nil {
		fmt.Fprintf(os.Stderr, "sqlx select client failed: %v\n", err)
		return
	}

	fmt.Println(cl)

	type Menu struct {
		Name        string      `db:"name"`
		Description pgtype.Text `db:"description"`
	}

	menu := make([]Menu, 0)
	if err := db.Select(&menu, "SELECT name, description FROM menu"); err != nil {
		fmt.Fprintf(os.Stderr, "sqlx select menu failed: %v\n", err)
		return
	}

	fmt.Println(menu)
}
