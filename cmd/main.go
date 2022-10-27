package main

import (
	"animescrapper/pkg/logger"
	"animescrapper/pkg/router"
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/joho/godotenv"
)

func main() {

	banner := figure.NewColorFigure("AniScrapper", "small", "cyan", true)

	if _, ok := os.LookupEnv("SERVER_PORT"); !ok {
		err := godotenv.Load("../.env")

		if err != nil {
			logger.Log.Fatal(err)
		}
	}

	banner.Print()
	fmt.Println("Created by NebulaZer0 & Provmawn")
	router.Run()
}
