package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "./styles")

	app.Get("/", MainPage)

	log.Fatal(app.Listen(":8080"))
}

func MainPage(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}
