package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"min-selhoz-backend/src/constants"
	"min-selhoz-backend/src/handlers"
	"min-selhoz-backend/src/repositories"
	"min-selhoz-backend/src/services"
	"os"
)

func NewDB() *sqlx.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Can't load .env file. Make sure it exist.")
	}

	dataSourceName := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", os.Getenv(constants.DbUser), os.Getenv(constants.DbPassword), os.Getenv(constants.DbName))
	db, err := sqlx.Connect("postgres", dataSourceName)
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
	tHeader.Get("/list/:tableId", tableHeaderHandler.ListByTableId())
	tHeader.Post("/create", tableHeaderHandler.Create())

	app.Listen(":3000")
}
