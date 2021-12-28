package main

import (
	"fmt"
	"log"
	"os"

	"gateway/configuration"
	"gateway/http"
	"gateway/middleware"

	"github.com/goccy/go-yaml"
)

func main() {
	file, err := os.Open("./conf.yaml")
	if err != nil {
		panic(err)
	}

	config := &configuration.Configuration{}
	dec := yaml.NewDecoder(file)
	err = dec.Decode(config)
	if err != nil {
		panic(err)
	}

	// Set Configuration
	http.Config(config)

	// Set Default Middleware
	http.Use(middleware.Recover())
	http.Use(middleware.Cors(config))

	// Listening
	fmt.Println("Listening...")
	log.Fatal(http.Run())
}
