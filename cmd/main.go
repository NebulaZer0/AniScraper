package main

import (
	"animescrapper/pkg/logger"
	"animescrapper/pkg/router"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/joho/godotenv"
)

func main() {

	p1 := figure.NewColorFigure("AniScrapper", "small", "cyan", true)

	if _, ok := os.LookupEnv("SERVER_PORT"); !ok {
		err := godotenv.Load("../.env")

		if err != nil {
			logger.Log.Fatal(err)
		}
	}

	p1.Print()
	router.Run()
}
