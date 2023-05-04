package services

import (
	"github.com/lclpedro/scaffold-golang-fiber/internal/scaffold/services/customer"
	"github.com/lclpedro/scaffold-golang-fiber/internal/scaffold/services/health"
	"github.com/lclpedro/scaffold-golang-fiber/pkg/mysql"
)

type AllServices struct {
	HealthService   health.Service
	CustomerService customer.Service
}

func NewAllServices(uow mysql.UnitOfWorkInterface) *AllServices {
	return &AllServices{
		HealthService:   health.NewHealthService(uow),
		CustomerService: customer.NewCustomerService(uow),
	}
}
