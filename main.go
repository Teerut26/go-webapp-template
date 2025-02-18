package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/teerut26/go-webapp-template/route"
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 1024 * 1024 * 20, // 10mb
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Static("/", "./web")

	apiGroup := app.Group("/api")

	apiGroup.Get("/hello", route.HelloHandler)

	log.Fatal(app.Listen(":3000"))
	log.Println("Server is running on port 3000")
}
