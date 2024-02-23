package server

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

func GracrfullShutdown(app *fiber.App) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		log.Println("server shutting down")
		_ = app.Shutdown()
	}()
}
