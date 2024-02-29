package service

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"sync"
	"time"
)

// Proxies is a global variable to store the proxies
var Proxies map[string]string
var Keys []string
var mutex sync.RWMutex

// UpdateProxies is a goroutine to update the proxies from Redis
func UpdateProxies(rdb *redis.Client) {
	log.Println("Starting the UpdateProxies goroutine")
	ctx := context.Background()
	for {
		// Get the proxies from Redis
		newProxies, err := rdb.HGetAll(ctx, "use_proxy").Result()
		if err != nil {
			log.Println("Error getting proxies from Redis:", err)
		} else {
			// Update the global Proxies variable
			mutex.Lock()
			Proxies = newProxies
			Keys = make([]string, 0, len(newProxies))
			for k := range newProxies {
				Keys = append(Keys, k)
			}
			mutex.Unlock()
		}
		log.Println("Proxies updated")

		// Sleep for 10 seconds
		time.Sleep(10 * time.Second)
	}
}
