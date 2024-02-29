package main

import (
	"github.com/go-redis/redis/v8"
	"log"
	"main/service"
	"net/http"
)

func main() {
	service.RegisterRoutes()
	config := LoadConfigFromEnviron()

	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisHost + ":" + config.RedisPort,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})

	go service.UpdateProxies(rdb)

	// startup the web server
	log.Println("Server is running on " + config.Host + ":" + config.Port)
	if err := http.ListenAndServe(config.Host+":"+config.Port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
