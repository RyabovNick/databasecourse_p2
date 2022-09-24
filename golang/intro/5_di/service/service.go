package service

import (
	"fmt"
)

type People struct {
	ID   string
	Name string
}

type storer interface {
	ListPeople() ([]People, error)
	GetPeopleByID(id string) (People, error)
}

type tax interface {
	GetTaxStatusByID(id string) (string, error)
}

type Service struct {
	Store storer
	Tax   tax
}

type PeopleWithTax struct {
	People
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
func (s *Service) GetPeopleByID(id string) ([]PeopleWithTax, error) {
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
