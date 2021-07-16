package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/rs/cors"
	"github.com/spinup-host/api"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", api.Hello)
	mux.HandleFunc("/createservice", api.CreateService)
	mux.HandleFunc("/githubAuth", api.GithubAuth)
	// TODO: remove http version
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"https://spinup.host", "http://spinup.host"},
	})
	err := http.ListenAndServe("localhost:3001", c.Handler(mux))
	if err != nil {
		log.Fatalf("FATAL: starting server %v", err)
	}
}
