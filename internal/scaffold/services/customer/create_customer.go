package customer

import (
	"context"
	"fmt"

	"github.com/lclpedro/scaffold-golang-fiber/internal/scaffold/domains"
	"github.com/lclpedro/scaffold-golang-fiber/internal/scaffold/repositories/address"
	"github.com/lclpedro/scaffold-golang-fiber/internal/scaffold/repositories/person"
	uow "github.com/lclpedro/scaffold-golang-fiber/pkg/unit_of_work"
)

type InputNewCustomer struct {
	Person struct {
		FirstName string
		LastName  string
		Age       int
	}
	Address struct {
		StreetName string
		ZipCode    string
		City       string
		State      string
		Country    string
	}
}

func (i *InputNewCustomer) ToDomains() (*domains.Person, *domains.Address) {
	personDomain := domains.NewPerson(i.Person.FirstName, i.Person.LastName, i.Person.Age, 0)

	addressDomain := domains.NewAddress(i.Address.StreetName, i.Address.ZipCode, i.Address.City, i.Address.State,
		i.Address.Country)

	return personDomain, addressDomain
}

func (s *customerService) getAddressRepository(ctx context.Context) (address.Repository, error) {
	repo, err := s.uow.GetRepository(ctx, "AddressRepository")
	if err != nil {
		return nil, err
	}
	return repo.(address.Repository), nil
}

func (s *customerService) getPersonRepository(ctx context.Context) (person.Repository, error) {
	repo, err := s.uow.GetRepository(ctx, "PersonRepository")
	if err != nil {
		return nil, err
	}
	return repo.(person.Repository), nil
}

func (s *customerService) CreateCustomer(ctx context.Context, input InputNewCustomer) error {
	return s.uow.Do(ctx, func(uow *uow.UnitOfWork) error {
		addressRepo, err := s.getAddressRepository(ctx)
		if err != nil {
			return nil
		}
		personRepo, err := s.getPersonRepository(ctx)
		if err != nil {
			return nil
		}

		personDomain, addressDomain := input.ToDomains()
		addressID, err := addressRepo.Save(addressDomain)
		if err != nil {
			return err
		}
		personDomain.AddressID = addressID
		personID, err := personRepo.Save(personDomain)
		if err != nil {
			return err
		}
		fmt.Println(fmt.Sprintf("Creste PersonID: %d", personID))
		return nil
	})
}
