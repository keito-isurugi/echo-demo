package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"echo-demo/internal/db"
	"echo-demo/internal/server"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("envを読み込めませんでした： %v", err)
	}

	dbConect, err := db.NewConect()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	
	router := server.SetupRouter(dbConect)

	router.Logger.Fatal(router.Start(":" + os.Getenv("SERVER_PORT")))
}
