package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	_ "cmAct/internal/models"
	"cmAct/internal/routes"
)

func main() {
	// Here you will need to set the waiting time for a response from the server,
	// the maximum number of requests, and if I remember something else
	engine := html.New("../templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	routes.RegisterRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
