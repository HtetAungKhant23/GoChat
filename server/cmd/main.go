package main

import (
	"fmt"
	"log"
	"server/db"
	"server/internal/user"
	"server/router"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %v", err)
	}

	userRepo := user.NewRepository(dbConn.GetDB())
	userSrv := user.NewService(userRepo)
	userHandler := user.NewHandler(userSrv)

	router.InitRouter(userHandler)
	router.Start("0.0.0.0:3001")

	fmt.Println("database successfully connected")
}
