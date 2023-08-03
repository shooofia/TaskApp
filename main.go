package main

import (
	"TaskApp/config"
	"TaskApp/route"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()
	config.ConnectDatabase()
	e := echo.New()
	e = route.InitRoute(e)
	e.Start(getPort())

}

func getPort() string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	return ":8080"
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}
