package proxy

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sakshamagrawal07/cache-proxy-server/cache"
	"github.com/sakshamagrawal07/cache-proxy-server/utils"
)

type ProxyObject struct {
	Origin string
	Cache  map[string]*cache.CacheObject
}

func NewProxy(origin string) *ProxyObject {
	return &ProxyObject{
		Origin: origin,
		Cache:  make(map[string]*cache.CacheObject), // Initialize the cache map
	}
}

func (p *ProxyObject) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	db := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:     "localhost:6379",
		Password: "", // Add password if needed
		DB:       0,
	})

	ctx := context.Background()

	result, err := db.Get(ctx, "name").Result()
	if err != nil {
		http.Error(w, "Error connecting to Redis", http.StatusInternalServerError)
		log.Printf("Error connecting to Redis: %v", err)
		return
	}

	fmt.Println("Result",result)

	CACHE_KEY := r.Method + ":" + r.URL.String()

	// Cache present
	if cache, ok := p.Cache[CACHE_KEY]; ok {
		utils.RespondWithHeaders(w, *cache.Response, cache.ResponseBody, "Hit", CACHE_KEY)
		return
	}

	originUrl := p.Origin + r.URL.String()
	res, err := http.Get(originUrl)

	if err != nil {
		http.Error(w, "Error in forwarding request.", http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		http.Error(w, "Error forwarding request body.", http.StatusInternalServerError)
		return
	}

	p.Cache[CACHE_KEY] = &cache.CacheObject{
		Response:     res,
		ResponseBody: body,
		Created:      time.Now(),
	}

	utils.RespondWithHeaders(w, *res, body, "MISS", CACHE_KEY)
}
