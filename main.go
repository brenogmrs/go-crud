package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading env file")
	}

	fmt.Println("NVIM NESSA PORRA", os.Getenv("TEST"))
}
