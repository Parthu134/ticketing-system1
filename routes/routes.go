package routes

import (
	"ticketing-system/controllers"
	"ticketing-system/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Post("/verify-otp", controllers.VerifyOTP)
	app.Post("/tickets", middleware.RequiredAuth, controllers.CreateTicket)
	app.Get("/tickets/tag/:tag", controllers.GetTicketsByTag)
	app.Get("/tickets/all", controllers.ListTicketsAdmins)
	app.Patch("/tickets/:id", middleware.RequiredAuth, controllers.UpdateTicketStatus)
	app.Get("/admin/subscription", controllers.SaveSubscription)
	app.Put("/tickets/:id/reply", middleware.AdminOnly, controllers.ReplyTicket)
	app.Delete("/tickets/:id", controllers.DeleteTicket)
}
