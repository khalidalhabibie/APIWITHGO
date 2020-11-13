package main

import (
    "github.com/joho/godotenv"
    "log"
    "fmt"
    
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  fmt.Println("Starts")

	log.Fatal(route.RunAPI(":8090"))
  
}