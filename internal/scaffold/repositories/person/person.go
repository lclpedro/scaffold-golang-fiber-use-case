package person

import (
	"fmt"

	"github.com/lclpedro/scaffold-golang-fiber/internal/scaffold/domains"
	"github.com/lclpedro/scaffold-golang-fiber/pkg/mysql"
)

type Repository interface {
	Save(person *domains.Person) (int64, error)
	// Find(personID int) (*domains.Person, error)
	// FindAll() ([]*domains.Person, error)
	// Delete() error
	// Update(personID int, person *domains.Person) error
}

type personRepository struct {
	dbConnection mysql.Connection ``
}

func NewPersonRepository(dbConnection mysql.Connection) Repository {
	return &personRepository{
		dbConnection: dbConnection,
	}
}

const insertPerson = `INSERT INTO person(first_name, last_name, age, address_id) VALUES (?,?,?,?)`

func (r *personRepository) Save(person *domains.Person) (int64, error) {
	query, err := r.dbConnection.Exec(insertPerson,
		person.FistName, person.LastName, person.Age, person.AddressID,
	)

	if err != nil {
		fmt.Println(fmt.Sprintf("PersonRepository - Error: %s", err.Error()))
		return 0, err
	}
	idInserted, err := query.LastInsertId()

	if err != nil {
		fmt.Println(fmt.Sprintf("PersonRepository - Error: %s", err.Error()))
		return 0, err
	}
	return idInserted, nil
}
