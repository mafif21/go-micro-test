package main

import (
	"github.com/joho/godotenv"
	"product/internal/app"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	app.StartApp()
}
