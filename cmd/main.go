package main

import (
	"animescrapper/pkg/logger"
	"animescrapper/pkg/router"

	"github.com/common-nighthawk/go-figure"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		logger.Log.Fatal(err)
	} else {
		title := figure.NewColorFigure("Anime Scrapper", "shadow", "green", true)
		title.Print()
		router.Run()
	}
}
