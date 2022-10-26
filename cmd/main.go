package main

import (
	"animescrapper/pkg/logger"
	"animescrapper/pkg/router"

	"github.com/common-nighthawk/go-figure"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	p1 := figure.NewColorFigure("AniScrapper", "small", "cyan", true)
	if err != nil {
		logger.Log.Fatal(err)
	} else {
		p1.Print()
		router.Run()
	}
}
