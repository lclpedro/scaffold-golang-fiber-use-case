package customer

import (
	"context"

	uow "github.com/lclpedro/scaffold-golang-fiber/pkg/unity_of_work"
)

type Service interface {
	CreateCustomer(ctx context.Context, input InputNewCustomer) error
}

type customerService struct {
	uow uow.UnityOfWorkInterface
}

func NewCustomerService(uow uow.UnityOfWorkInterface) Service {
	return &customerService{
		uow: uow,
	}
}
