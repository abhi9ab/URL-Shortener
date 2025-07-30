package database

import (
	// Required for Redis operations e.g. timeout, cancellation
	"context"
	// Redis client library
	"github.com/go-redis/redis/v8"
	// For reading .env variables using os.Getenv
	"os"
)
// Ctx is a global context passed into every Redis command (e.g., client.Get(Ctx, key)).
// In a real-world app, you'd often use context.WithTimeout() for timeout handling, but context.Background() is fine for most simple CLI/web use cases.
var Ctx = context.Background()

// This function returns a pointer to a Redis client, so you can reuse the connection across your app.
// Addr: Redis server address (e.g. localhost:6379 or your AWS ElastiCache endpoint).
// Password: If your Redis requires a password (leave empty if not).
// DB: Redis has multiple numbered logical databases (0 to 15), useful for separating data logically. You're allowing the caller to choose.
func CreateClient(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
		DB:       dbNo,
	})
	return rdb
}
