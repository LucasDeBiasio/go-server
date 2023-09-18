package main

import (
	"fmt"
	"log"
	"os"
)


func main() {
	fmt.Println("hello world")

	godotenv.Load(".env")
	// Getenv is a os function from the imported "os" library that gets the environment details from the .env file
	portString := os.Getenv("PORT")
	if (portString == "") {
		// log.Fatal ends program with error code 1 and returns the error message
		log.Fatal("PORT is not found in the environment")
	}
	fmt.Println("PORT:", portString)
	
}