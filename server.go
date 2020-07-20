package main

import (
	"github.com/captainblue210/goliff/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

func init() {
	// 環境変数を読み込む
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}
	// Info以上のレベルを出力する設定
	logrus.SetLevel(logrus.DebugLevel)
	// フォーマッターを指定
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	e := echo.New()

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	// Routes
	routes.Init(e)

	// start server
	e.Logger.Fatal(e.Start(":8080"))
}
