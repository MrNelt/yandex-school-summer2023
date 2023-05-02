package routes

import "github.com/labstack/echo/v4"

func SetupRoutes(e *echo.Echo) {
	e.GET("/ping", ping)
	e.POST("/couriers", createCouriers)
}
