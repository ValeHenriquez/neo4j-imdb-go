package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ValeHenriquez/neo4j-imdb-go/database"
	"github.com/ValeHenriquez/neo4j-imdb-go/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Server starting")
	godotenv.Load()
	fmt.Println("Loaded env variables")

	fmt.Println("Connecting to database")
	database.Setup()
	defer database.Close()
	fmt.Println("Database connected")

	app := fiber.New()
	app.Use(cors.New())

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not set in .env file")
	}

	routes.Setup(app)
	fmt.Println("Server started, listening on port " + portString + "...")
	app.Listen(":" + portString)

}
