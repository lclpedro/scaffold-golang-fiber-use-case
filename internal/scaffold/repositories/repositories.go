package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/lclpedro/scaffold-golang-fiber/internal/scaffold/repositories/address"
	"github.com/lclpedro/scaffold-golang-fiber/internal/scaffold/repositories/health"
	"github.com/lclpedro/scaffold-golang-fiber/internal/scaffold/repositories/person"
	"github.com/lclpedro/scaffold-golang-fiber/pkg/mysql"
)

type AllRepositories struct {
	HealthRepository health.Repository
	Person           person.Repository
	Address          address.Repository
}

func RegistryRepositories(uow mysql.UnitOfWorkInterface, dbConnection mysql.Connection) mysql.UnitOfWorkInterface {
	uow.Register("HealthRepository", func(tx *sqlx.Tx) interface{} {
		repo := health.NewHealthRepository(dbConnection)
		return repo
	})

	uow.Register("PersonRepository", func(tx *sqlx.Tx) interface{} {
		repo := person.NewPersonRepository(dbConnection)
		return repo
	})

	uow.Register("AddressRepository", func(tx *sqlx.Tx) interface{} {
		repo := address.NewAddressRepository(dbConnection)
		return repo
	})
	return uow
}
