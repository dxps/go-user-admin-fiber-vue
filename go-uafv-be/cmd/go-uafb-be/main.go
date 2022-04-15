package main

import (
	"dxps.io/go-user-admin-fiber-vue/be/internal/database"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func main() {

	db, dbCloseFn, err := database.Connect()
	if err != nil {
		log.Panicf("Cound not connect to the database: %v", err.Error())
	}
	defer func() {
		log.Warnf("Database conn closed err: %v", dbCloseFn())
	}()
	log.Debugf("Database conn established: %v", db)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	log.Warnf("HTTP Listening err: %v", app.Listen(":3000"))

}
