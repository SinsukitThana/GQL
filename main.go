package main

import (
	"context"
	"os"

	"github.com/SinsukitThana/GQL/database"
	"github.com/SinsukitThana/GQL/model/person"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	app := fiber.New()

	//Load .env
	env := os.Getenv("MODE")
	if "" == env {
		env = "dev"
	}

	if "prod" == env {
		godotenv.Load(".env.prod")
	}
	godotenv.Load(".env." + env)

	db := database.SetDatabase()
	app.Get("/", func(c *fiber.Ctx) error {
		person := new(person.Persons)

		if err := db.NewSelect().Model(person).Limit(1).Scan(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return c.JSON(person)
	})

	app.Listen(":3000")
}
