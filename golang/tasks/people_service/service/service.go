package service

import (
	"fmt"

	"github.com/RyabovNick/databasecourse_2/golang/tasks/people_service/service/store"
)

// TODO:
// 1. Посмотреть видео SQL in GO https://www.youtube.com/watch?v=INn0jtSOMco&list=PL4wpCvLmqD_TKIAmEmu4piCGWJFvzxXKh&index=5
// 2. Реализовать store.go (бес использования sqlx): ListPeople, GetPeopleByID
// 3. Познакомиться с https://github.com/golang-migrate/migrate
// 4. И написать в store функцию, которая будет выполнять миграции при запуске приложения

type storer interface {
	ListPeople() ([]store.People, error)
	GetPeopleByID(id int) (store.People, error)
}

type tax interface {
	GetTaxStatusByID(id int) (string, error)
}

type Service struct {
	Store storer
	Tax   tax
}

type PeopleWithTax struct {
	store.People
	TaxStatus string
}

// ListPeople gets list of people from store and
// add tax status from tax service
func (s *Service) ListPeople() ([]PeopleWithTax, error) {
	list, err := s.Store.ListPeople()
	if err != nil {
		return nil, fmt.Errorf("list people: %w", err)
	}

	people := make([]PeopleWithTax, 0, len(list))
	for _, l := range list {
		st, err := s.Tax.GetTaxStatusByID(l.ID)
		if err != nil {
			return nil, fmt.Errorf("get tax status: %w", err)
		}

		people = append(people, PeopleWithTax{
			People:    l,
			TaxStatus: st,
		})
	}

	return people, nil
}

// GetPeopleByID gets people by id from store and
// add tax status from tax service
func (s *Service) GetPeopleByID(id int) ([]PeopleWithTax, error) {
	p, err := s.Store.GetPeopleByID(id)
	if err != nil {
		return nil, fmt.Errorf("get people by id: %w", err)
	}

	st, err := s.Tax.GetTaxStatusByID(p.ID)
	if err != nil {
		return nil, fmt.Errorf("get tax status: %w", err)
	}

	return []PeopleWithTax{{
		People:    p,
		TaxStatus: st,
	}}, nil
}
