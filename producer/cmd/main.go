package main

import (
	"producer/internal/app"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app.Start()
}
