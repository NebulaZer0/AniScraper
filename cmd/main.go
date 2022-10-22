package main

import (
	"animescrapper/pkg/logger"
	"animescrapper/pkg/router"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		logger.Log.Fatal(err)
	} else {
		router.Run()
	}
}
