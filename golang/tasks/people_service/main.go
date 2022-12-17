package main

import (
	"fmt"

	"github.com/RyabovNick/databasecourse_2/golang/tasks/people_service/service"
	"github.com/RyabovNick/databasecourse_2/golang/tasks/people_service/service/store"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db := store.NewStore("")

	serv := service.Service{
		Store: db,
		Tax:   &t{},
	}

	fmt.Println(serv.GetPeopleByID(1)) // example
}

// simple fake tax realization
type t struct{}

func (t *t) GetTaxStatusByID(id int) (string, error) {
	return "", nil
}
