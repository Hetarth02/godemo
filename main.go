package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hello World!")
	app := fiber.New()
	app.Get("/", hello)
	log.Fatal(app.Listen(":3000"))
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}
