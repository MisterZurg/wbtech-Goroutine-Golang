package main

import (
	"github.com/MisterZurg/wbtech-Goroutine-Golang/l0/internal/handlers"
	"github.com/MisterZurg/wbtech-Goroutine-Golang/l0/internal/repository"
	"github.com/gofiber/fiber/v3"
)

func main() {
	//cfg, err := config.New()
	//if err != nil {
	//
	//}
	//repo, err := repository.New(cfg.GetPostgresConnectionString())
	//if err != nil {
	//
	//}
	repo := repository.Repository{}
	svc := handlers.New(repo)

	app := fiber.New()
	app.Get("/orders/:order_uid", svc.GetByID).
		Post("/orders", svc.Create)

	app.Listen(":3000")
}
