package main

// standard library for printing, logging errors, and accessing environment varibles
import (
	"fmt"
	"log"
	"os"

	"github.com/abhi9ab/URL-Shortener/routes"
	// Web framework like Express.
	"github.com/gofiber/fiber/v2"
	// Middleware to log HTTP requests.
	"github.com/gofiber/fiber/v2/middleware/logger"
	// Loads .env file like dotenv in Node.js
	"github.com/joho/godotenv"
)

// registers the routes
func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	// Initializes the Fiber web server like app := express() in Node
	app := fiber.New()

	// app.Use(csrf.New())
	// Adds logging middleware (like morgan in Express) to log each incoming HTTP request.
	app.Use(logger.New())

	// Registers your custom routes.
	setupRoutes(app)

	// Starts the server on the port defined in .env, e.g., APP_PORT=:3000.
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
