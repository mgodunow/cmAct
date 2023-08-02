package routes

import (
	"cmAct/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

var RegisterRoutes = func(app *fiber.App) {
	// authGroup := app.Group("/", middleware.UserIdentity)
	// authGroup.Get("/home", controllers.Home)
	app.Get("/home", controllers.Home)
	app.Static("/", "../static")
	app.Post("/login", controllers.Login)
	app.Post("/registration", controllers.Register)
	// Need to think about how control groups
	// app.Get("/profile", controllers.Profile)
	// app.Get("/act", controllers.Act)
}
