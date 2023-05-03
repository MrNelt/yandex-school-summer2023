package routes

import (
	"github.com/labstack/echo/v4"
	"yandex-team.ru/bstask/routes/courier-conroller"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/ping", ping)

	e.POST("/couriers", courier_conroller.CreateCouriers)
	e.GET("/couriers/:id", courier_conroller.GetCourierByID)
	e.GET("/couriers", courier_conroller.GetCouriers)
	e.DELETE("/couriers/:id", courier_conroller.DeleteCourierByID)
}
