package utils

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func RespondWithHeaders(w http.ResponseWriter, statusCode int, body []byte, cacheHeader string, cacheKey string) {
	w.Header().Set("X-Cache", cacheHeader)
	w.WriteHeader(statusCode)
	w.Write(body)
}

func ClearCache() {
	// Connect to Redis
	db := redis.NewClient(&redis.Options{
		Addr:     "localhost:637",
		Password: "",
		DB:       0,
	})

	// Clear the cache
	err := db.FlushDB(context.Background()).Err()

	if err != nil {
		fmt.Println("Error clearing cache")
		log.Fatal(err)
	} else {
		fmt.Println("Cache cleared successfully")
	}
}
