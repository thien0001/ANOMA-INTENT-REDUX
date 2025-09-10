package main

import (
	"anoma_intent_app/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// static frontend
	app.Static("/", "./static")

	// API
	app.Post("/intent", handlers.IntentHandler)
	app.Get("/history", handlers.HistoryHandler)

	log.Println("🚀 Server running at http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
