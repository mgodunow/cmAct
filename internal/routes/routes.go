package routes

import (
	"cmAct/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

var RegisterRoutes = func(app *fiber.App) {
	app.Static("/", "../templates", fiber.Static{
		Index: "index.html",
	})
	// app.Post("/login", controllers.Login)
	app.Post("/registration", controllers.Register)
	// app.Get("/profile", controllers.Profile)
	// app.Get("/act", controllers.Act)
}
