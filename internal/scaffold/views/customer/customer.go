package customer

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lclpedro/scaffold-golang-fiber/internal/scaffold/services/customer"
)

type View interface {
	CreateCustomerHandler(c *fiber.Ctx) error
}

type customerView struct {
	CustomerService customer.Service
}

func NewCustomerView(customerService customer.Service) View {
	return &customerView{
		CustomerService: customerService,
	}
}
