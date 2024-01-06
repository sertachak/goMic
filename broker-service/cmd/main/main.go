package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct {
}

type Route struct {
	Config
	routeName string
}

func NewConfig() *Config {
	return &Config{}
}

func main() {
	app := Route{}
	log.Printf("/T", app)
	routes := app.routes1()

	log.Printf("Starting broker service on port %s\n", webPort)

	//define https server
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: routes,
	}

	//start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
