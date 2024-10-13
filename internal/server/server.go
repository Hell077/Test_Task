package server

import (
	"github.com/Hell077/Test_Task/internal/api"
	"github.com/gofiber/fiber/v2"
	"log"
)

func Server() {
	app := fiber.New()
	api.Api(app)
	log.Fatal(app.Listen(":3000"))
}
