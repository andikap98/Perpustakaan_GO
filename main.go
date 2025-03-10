package main

import (
	"github.com/gofiber/fiber/v2"
	"shellrean.id/belajar-golang-rest-api/internal/api"
	"shellrean.id/belajar-golang-rest-api/internal/config"
	"shellrean.id/belajar-golang-rest-api/internal/connection"
	"shellrean.id/belajar-golang-rest-api/internal/repository"
	"shellrean.id/belajar-golang-rest-api/internal/service"
)

func main() {
	cnf := config.Get()

	dbConnection := connection.GetDatabase(cnf.Database)

	app := fiber.New()

	customerRepository := repository.NewCustomer(dbConnection)

	customerService := service.NewCustomer(customerRepository)

	api.NewCustomer(app, customerService)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
