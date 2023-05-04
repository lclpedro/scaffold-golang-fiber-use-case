package customer

import (
	"context"

	"github.com/lclpedro/scaffold-golang-fiber/pkg/mysql"
)

type Service interface {
	CreateCustomer(ctx context.Context, input InputNewCustomer) error
}

type customerService struct {
	uow mysql.UnitOfWorkInterface
}

func NewCustomerService(uow mysql.UnitOfWorkInterface) Service {
	return &customerService{
		uow: uow,
	}
}
