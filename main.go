package main

import (
	"log"
	"os"
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
	port:=os.Getenv("PORT")
	if port==""{
		port="2500"
	}
	log.Printf("Server running on port %s",port)
	app.Listen("0.0.0.0:"+port)
}
