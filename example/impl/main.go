package main

import (
	greeterpb "example/gen/go/greeter"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func main() {

	app := fiber.New()

	greeterpb.RegisterGreeterServiceFiberRoutes(app, &service{}, nil)

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal().Err(err).Msg("start app")
	}
}
