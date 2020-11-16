package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"

	"github.com/khalidalhabibie/APIWITHGO/route"

)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Starts")
	log.Fatal(route.Run(":8080"))

}
