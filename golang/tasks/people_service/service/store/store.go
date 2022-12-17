package store

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type Store struct {
	conn *pgx.Conn
}

type People struct {
	ID   int
	Name string
}

// NewStore creates new database connection
func NewStore(connString string) *Store {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		panic(err)
	}

	// make migration

	return &Store{
		conn: conn,
	}
}

func (s *Store) ListPeople() ([]People, error) {
	return nil, nil
}

func (s *Store) GetPeopleByID(id int) (People, error) {
	return People{}, nil
}
