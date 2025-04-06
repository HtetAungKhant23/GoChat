package main

import (
	"fmt"
	"log"
	"server/db"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	_, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %v", err)
	}

	fmt.Println("database successfully connected")
}
