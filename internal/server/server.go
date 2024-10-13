package server

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"test/internal/api"
)

func Server() {
	app := fiber.New()
	api.Api(app)
	log.Fatal(app.Listen(":3000"))
}
