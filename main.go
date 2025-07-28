package main

import (
	"ticketing-system/config"
	"ticketing-system/routes"
	"ticketing-system/services"

	"github.com/gofiber/fiber/v2"
	"github.com/robfig/cron/v3"
)

func main() {
	config.ConnectDB()
	app := fiber.New()
	routes.Setup(app)

	c := cron.New()
	c.AddFunc("@every 1h", services.AutoUpdatingStatus)
	c.Start()
	
	app.Listen(":2500")
}
