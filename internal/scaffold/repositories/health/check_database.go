package health

import (
	"github.com/lclpedro/scaffold-golang-fiber/pkg/mysql"
)

type healthRepository struct {
	dbConnection mysql.Connection
}

func (h *healthRepository) GetDatabaseCheck() error {
	query := h.dbConnection.QueryRow("SELECT 1;")
	if err := query.Err(); err != nil {
		return err
	}
	return nil
}
