package proxy

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
	"github.com/sakshamagrawal07/cache-proxy-server/utils"
)

type ProxyObject struct {
	Origin string
}

// NewProxy creates a new ProxyObject
func NewProxy(origin string) *ProxyObject {
	return &ProxyObject{
		Origin: origin,
	}
}

func (p *ProxyObject) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	// Connect to Redis
	db := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	CACHE_KEY := r.Method + ":" + r.URL.String() + ":" + "body"

	// Get the response body from Redis
	body, err := db.Get(ctx, CACHE_KEY).Result()

	if err == redis.Nil {

		// Cache not present
		//Contruct the origin URL
		var originUrl string

		// If the request is for the root path, then the origin URL is the origin itself
		if r.URL.String() == "/" {
			originUrl = p.Origin
		} else {
			originUrl = p.Origin + r.URL.String()
		}

		// Forward the request to the origin server
		res, err := http.Get(originUrl)
		if err != nil {
			http.Error(w, "Error in forwarding request.", http.StatusInternalServerError)
			return
		}

		defer res.Body.Close()

		// Read the response body
		body, err := io.ReadAll(res.Body)
		if err != nil {
			http.Error(w, "Error forwarding request body.", http.StatusInternalServerError)
			return
		}

		// Cache the response body
		// Store the response body in Redis
		bodyCacheError := db.Set(ctx, CACHE_KEY, body, 0).Err()
		if bodyCacheError != nil {
			fmt.Println("Error caching response body")
			log.Fatal(bodyCacheError)
		}

		fmt.Println("Cache Miss")
		fmt.Println("Cache Key", CACHE_KEY)
		fmt.Println("Origin URL", originUrl)
		fmt.Println("Response", res)
		fmt.Println("Body", body)
		fmt.Printf("Response type %T\n", res)
		fmt.Printf("Body type %T\n", body)

		// Respond to the client
		utils.RespondWithHeaders(w, res.StatusCode, body, "MISS", CACHE_KEY)
	} else if err != nil {
		log.Fatal(err)
		return
	} else {
		fmt.Println("Cache Hit")
		fmt.Println("Body", body)
		utils.RespondWithHeaders(w, http.StatusOK, []byte(body), "Hit", CACHE_KEY)
	}
}
