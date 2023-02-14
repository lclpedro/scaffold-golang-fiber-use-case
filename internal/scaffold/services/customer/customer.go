package customer

import (
	"context"

	uow "github.com/lclpedro/scaffold-golang-fiber/pkg/unit_of_work"
)

type Service interface {
	CreateCustomer(ctx context.Context, input InputNewCustomer) error
}

type customerService struct {
	uow uow.UnitOfWorkInterface
}

func NewCustomerService(uow uow.UnitOfWorkInterface) Service {
	return &customerService{
		uow: uow,
	}
}
