package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"min-selhoz-backend/src/handlers"
	"min-selhoz-backend/src/repositories"
	"min-selhoz-backend/src/services"
)

func NewDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "")
	if err != nil {
		panic(err.Error())
	}

	return db
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
	tHeader.Post("/create", tableHeaderHandler.Create())

	app.Listen(":3000")
}
