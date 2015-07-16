package main

import (
	"net/http"
	"log"

	"github.com/cyarie/linksecret/application/router"
	"github.com/cyarie/linksecret/application/environment"
	"gopkg.in/redis.v3"
)

func main() {
	// Setup our redis connection
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.Println("Error connecting to redis.")
	} else {
		log.Println(pong)
	}

	env := &environment.Env{
		CL: client,
	}

	router := router.Router(env)

	http.ListenAndServe(":8080", router)
}