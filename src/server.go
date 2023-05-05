package main

import (
	"github.com/labstack/echo/v4"
	_ "github.com/swaggo/echo-swagger/example/docs"
	"yandex-team.ru/bstask/routes"
)

// @title Yandex Lavka
// @version 1.0
// @description API Server for Yandex Lavka

// @host localhost:8000
// @BasePath /

func main() {
	e := setupServer()
	e.Logger.Fatal(e.Start(":8080"))
}

func setupServer() *echo.Echo {
	e := echo.New()
	routes.SetupRoutes(e)
	return e
}
