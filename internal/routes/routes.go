package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mgodunow/cmAct/internal/controllers"
)

var RegisterRoutes = func(c *fiber.Ctx, app *fiber.App) {
	app.Get("/", controllers.MainPage)
	app.Get("/login", controllers.LoginPage)
	app.Get("/reg", controllers.RegisterPage)
	app.Get("/profile", controllers.ProfilePage)
	app.Get("/act", controllers.ActPage)

}
