package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"min-selhoz-backend/src/handlers"
	"min-selhoz-backend/src/repositories"
	"min-selhoz-backend/src/services"
)

func NewDB() *sqlx.DB {
	return nil
}

func main() {
	app := fiber.New()

	db := NewDB()

	tableHeaderRepo := repositories.NewTableHeaderRepository(db)
	tableHeaderService := services.NewTableHeaderService(tableHeaderRepo)
	tableHeaderHandler := handlers.NewTableHeaderHandler(tableHeaderService)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api")

	tHeader := api.Group("/table_header")
	tHeader.Get("/list/:table_id", tableHeaderHandler.ListByTableId())
	tHeader.Post("/list/create", tableHeaderHandler.Create())

	app.Listen(":3000")
}
