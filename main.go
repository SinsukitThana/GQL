package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"

	"github.com/SinsukitThana/GQL/database"
	"github.com/SinsukitThana/GQL/model/person"
	"github.com/SinsukitThana/GQL/model/workflow"
	"github.com/SinsukitThana/GQL/service/datafromgql"
	"github.com/SinsukitThana/GQL/service/server"
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

	var Workflow []workflow.Workflow

	app.Get("/GetDataGQL", func(c *fiber.Ctx) error {
		m := c.Queries()
		getQuery := m["query"]
		person := new([]person.Persons)
		fmt.Println(env)

		if err := db.NewSelect().Model(person).Scan(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		data := datafromgql.DataFromGQL(db, *person, getQuery)
		textData := fmt.Sprintf("%s", data)
		return c.JSON(textData)
	})

	app.Post("/GetWorkflow", func(c *fiber.Ctx) error {

		getQuery := fmt.Sprintf("%s", c.BodyRaw())

		randnumber, _ := rand.Int(rand.Reader, big.NewInt(100))
		newWf := workflow.Workflow{WorkflowID: fmt.Sprintf("%s", randnumber), WorkflowName: fmt.Sprintf("wf %s", randnumber)}

		Workflow = append(Workflow, newWf)

		data := datafromgql.DataWorkFlowFromGQL(db, Workflow, getQuery)
		textData := fmt.Sprintf("%s", data)
		return c.JSON(textData)
	})
	server.GracrfullShutdown(app)
	app.Listen(":3000")
}
