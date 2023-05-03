package main

import (
	"github.com/labstack/echo/v4"
	"yandex-team.ru/bstask/routes"
	"yandex-team.ru/bstask/storage"
)

func init() {
	storage.Init()
}

func main() {
	e := setupServer()
	e.Logger.Fatal(e.Start(":8080"))
}

func setupServer() *echo.Echo {
	e := echo.New()
	routes.SetupRoutes(e)
	return e
}
