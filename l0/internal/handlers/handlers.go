package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
)

type Repository interface {
	CreateOrder(order Order)
	GetOrder(string) Order
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{repo: repo}
}

// GetByID — GET /orders/:order_uid
func (svc *Service) GetByID(c fiber.Ctx) error {
	orderUID := c.Params("order_uid")
	return c.SendString(orderUID)
}

// Create — POST /orders
func (svc *Service) Create(c fiber.Ctx) error {
	p := new(Order)

	if err := c.Bind().Body(p); err != nil {
		return err
	}

	return c.SendString(fmt.Sprintf("%+v", p))
}
