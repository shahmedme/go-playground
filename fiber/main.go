package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Static("/", "./public")

	app.Get("/", HomeHandler)
	app.Get("/about/:name", AboutHandler)
	app.Get("/contact", ContactHandler)

	app.Listen(":8000")
}

func HomeHandler(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

func AboutHandler(c *fiber.Ctx) error {
	return c.SendString(c.Params("name"))
}

func ContactHandler(c *fiber.Ctx) error {
	return c.SendString("this is about page")
}
