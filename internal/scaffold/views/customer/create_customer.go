package customer

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/lclpedro/scaffold-golang-fiber/internal/scaffold/services/customer"
)

type NewCustomerRequest struct {
	Person struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Age       int    `json:"age"`
	} `json:"person"`
	Address struct {
		StreetName string `json:"street_name"`
		ZipCode    string `json:"zip_code"`
		City       string `json:"city"`
		State      string `json:"state"`
		Country    string `json:"country"`
	} `json:"address"`
}

func (v *customerView) CreateCustomerHandler(c *fiber.Ctx) error {
	ctx := context.Background()
	customerBody := new(NewCustomerRequest)
	if err := c.BodyParser(customerBody); err != nil {
		return err
	}
	input := customer.InputNewCustomer{
		Person: struct {
			FirstName string
			LastName  string
			Age       int
		}{
			FirstName: customerBody.Person.FirstName,
			LastName:  customerBody.Person.LastName,
			Age:       customerBody.Person.Age,
		},
		Address: struct {
			StreetName string
			ZipCode    string
			City       string
			State      string
			Country    string
		}{
			StreetName: customerBody.Address.StreetName,
			ZipCode:    customerBody.Address.ZipCode,
			City:       customerBody.Address.City,
			State:      customerBody.Address.State,
			Country:    customerBody.Address.Country,
		},
	}

	err := v.CustomerService.CreateCustomer(ctx, input)
	message := struct {
		Msg string `json:"msg"`
	}{}
	if err != nil {
		message.Msg = "Error in create new customer"
		c.Status(http.StatusBadRequest)
		return c.JSON(message)
	}
	message.Msg = "Customer created successfully"
	c.Status(http.StatusCreated)
	return c.JSON(message)
}
