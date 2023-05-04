package repositories

import (
	"database/sql"

	"github.com/lclpedro/scaffold-golang-fiber/internal/scaffold/repositories/address"
	"github.com/lclpedro/scaffold-golang-fiber/internal/scaffold/repositories/health"
	"github.com/lclpedro/scaffold-golang-fiber/internal/scaffold/repositories/person"
	"github.com/lclpedro/scaffold-golang-fiber/pkg/mysql"
	uow "github.com/lclpedro/scaffold-golang-fiber/pkg/unit_of_work"
)

type AllRepositories struct {
	HealthRepository health.Repository
	Person           person.Repository
	Address          address.Repository
}

func RegistryRepositories(uow uow.UnitOfWorkInterface, dbConnection mysql.Connection) uow.UnitOfWorkInterface {
	uow.Register("HealthRepository", func(tx *sql.Tx) interface{} {
		repo := health.NewHealthRepository(dbConnection)
		return repo
	})

	uow.Register("PersonRepository", func(tx *sql.Tx) interface{} {
		repo := person.NewPersonRepository(dbConnection)
		return repo
	})

	uow.Register("AddressRepository", func(tx *sql.Tx) interface{} {
		repo := address.NewAddressRepository(dbConnection)
		return repo
	})
	return uow
}
