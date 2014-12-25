package main

import (
	"flag"
	"fmt"
)

var CONFIG Config

func main() {
	// Parsing command line argument
	env := flag.String("env", "development", "the environment we are running")
	flag.Parse()

	// Load configuration from config file
	CONFIG = Config{}
	err := CONFIG.LoadConfig(*env)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Show configuration
	fmt.Printf("ENVIRONMENT: %s \n", *env)
	fmt.Printf("PORT: %s \n", CONFIG.Port)
	fmt.Printf("DATABASE: %s \n", CONFIG.Database)
}
