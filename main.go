package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Server starting")
	godotenv.Load()
	fmt.Println("Loaded env variables")

}
