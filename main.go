package main

import (
	"fmt"
	"log"
	"os"

	"github.com/goccy/go-yaml"
	"gitlab.com/zaen/gateway/configuration"
	"gitlab.com/zaen/gateway/http"
	"gitlab.com/zaen/gateway/middleware"
)

func main() {
	file, err := os.Open("./conf.yaml")
	if err != nil {
		panic(err)
	}

	var config configuration.Configuration
	dec := yaml.NewDecoder(file)
	err = dec.Decode(&config)
	if err != nil {
		return
	}

	// Set Configuration
	http.Config(&config)

	// TODO: Set Middleware
	http.Use(middleware.Recover())

	// Listening
	fmt.Println("Listening...")
	log.Fatal(http.Run())
}
