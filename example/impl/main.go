package main

import (
	greeterpb "example/gen/go/greeter"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	greeterpb.RegisterGreeterServiceFiberRoutes(app, &service{}, nil)

	app.Listen(":8080")
}
