package address

import (
	"fmt"
	"github.com/lclpedro/scaffold-golang-fiber/internal/scaffold/domains"
	"github.com/lclpedro/scaffold-golang-fiber/pkg/mysql"
)

type Repository interface {
	Save(address *domains.Address) (int64, error)
	// Find(addressID int) (*domains.Address, error)
	// FindAll() ([]*domains.Address, error)
	// Delete() error
	// Update(addressID int, address *domains.Address) error
}

type addressRepository struct {
	dbConnection mysql.Connection
}

func NewAddressRepository(dbConnection mysql.Connection) Repository {
	return &addressRepository{
		dbConnection: dbConnection,
	}
}

const insertAddress = `INSERT INTO address(street_name, zip_code, city, state, country) VALUES (?,?,?,?,?)`

func (r *addressRepository) Save(address *domains.Address) (int64, error) {
	query, err := r.dbConnection.Exec(insertAddress,
		address.StreetName, address.ZipCode, address.City, address.State, address.Country,
	)

	if err != nil {
		fmt.Println(fmt.Sprintf("AddressRepository - Error: %s", err.Error()))
		return 0, err
	}
	idInserted, err := query.LastInsertId()

	if err != nil {
		fmt.Println(fmt.Sprintf("AddressRepository - Error: %s", err.Error()))
		return 0, err
	}
	return idInserted, nil
}
