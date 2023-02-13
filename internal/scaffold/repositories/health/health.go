package health

import (
	"github.com/lclpedro/scaffold-golang-fiber/pkg/mysql"
)

type Repository interface {
	GetDatabaseCheck() error
}

func NewHealthRepository(dbConnection mysql.Connection) Repository {
	return &healthRepository{
		dbConnection: dbConnection,
	}
}
