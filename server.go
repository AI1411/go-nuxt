package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"go-nuxt-youtube/middlewares"
	"go-nuxt-youtube/routes"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env")
	}
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	e := echo.New()

	//middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middlewares.YouTubeService())
	e.Use(middlewares.DatabaseService())
	e.Use(middlewares.Firebase())

	//routes
	routes.Init(e)

	//サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}
