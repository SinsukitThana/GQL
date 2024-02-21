package main

import (
	"context"
	"fmt"
	"os"

	"github.com/SinsukitThana/GQL/database"
	"github.com/SinsukitThana/GQL/model/person"
	"github.com/SinsukitThana/GQL/service/datafromgql"
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

	app.Get("/GetDataGQL", func(c *fiber.Ctx) error {
		person := new([]person.Persons)
		fmt.Println(env)
		if err := db.NewSelect().Model(person).Scan(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		data := datafromgql.DataFromGQL(db, *person)
		textData := fmt.Sprintf("%s \n", data)
		return c.JSON(textData)
	})

	app.Listen(":3000")
}
