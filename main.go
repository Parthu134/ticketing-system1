package main

import (
	"ticketing-system/config"
	"ticketing-system/routes"
	"ticketing-system/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/robfig/cron/v3"
)

func main() {
	config.ConnectDB()
	app := fiber.New()
	routes.Setup(app)

	c := cron.New()
	go func() {
		for {
			services.CheckOldTicketsAndNotify()
			time.Sleep(5 * time.Hour)
		}
	}()
	c.Start()
	app.Listen(":2500")
}
