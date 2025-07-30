package routes

import (
	"github.com/abhi9ab/URL-Shortener/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

// ResolveURL ...
func ResolveURL(c *fiber.Ctx) error {
	// get the short from the url
	// Gets the path param from /:url — e.g., if path is /abc123, url = "abc123".
	url := c.Params("url")
	// query the db to find the original URL, if a match is found
	// increment the redirect counter and redirect to the original URL
	// else return error message
	
	// You create a Redis client for logical DB 0 — assumed to store URL mappings.
	r := database.CreateClient(0)
	// defer r.Close() ensures connection is closed after function exits (like finally).
	defer r.Close()

	// If the key is not found, Redis returns redis.Nil.
	// If some other error (e.g., Redis server down), return 500.
	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "short not found on database",
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot connect to DB",
		})
	}
	// increment the counter
	// Connects to Redis DB 1 to keep analytics data separate.
	rInr := database.CreateClient(1)
	defer rInr.Close()

	// ignoring the error here with _ = ... — that’s fine for optional logging.
	_ = rInr.Incr(database.Ctx, "counter")

	// redirect to original URL
	// Sends an HTTP 301 Moved Permanently response, telling the browser to go to the original URL.
	// You could use 302 instead if you want temporary redirect.
	return c.Redirect(value, 301)
}
