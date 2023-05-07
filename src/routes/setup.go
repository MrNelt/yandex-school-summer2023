package routes

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"os"
	"yandex-team.ru/bstask/routes/courier_controller"
	"yandex-team.ru/bstask/routes/order_controller"
	"yandex-team.ru/bstask/storage"
)

func SetupRoutes(e *echo.Echo) {
	dsl := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_MODELESS"))
	pg, err := storage.NewPgStorage(dsl)
	if err != nil {
		log.Panic(err)
	}

	e.GET("/ping", ping)

	courierController := courier_controller.NewCourierControllerHandler(pg)
	e.POST("/couriers", courierController.CreateCouriers)
	e.GET("/couriers/:id", courierController.GetCourierByID)
	e.GET("/couriers", courierController.GetCouriers)
	e.DELETE("/couriers/:id", courierController.DeleteCourierByID)
	e.GET("/couriers/meta-info/:id", courierController.CalculateRatingCourier)

	orderController := order_controller.NewOrderControllerHandler(pg)
	e.POST("/orders", orderController.CreateOrders)
	e.GET("/orders/:id", orderController.GetOrderByID)
	e.GET("/orders", orderController.GetOrders)
	e.DELETE("/orders/:id", orderController.DeleteOrderByID)
	e.POST("/orders/complete", orderController.CompleteOrder)
	e.GET("/orders/complete", orderController.GetOrdersComplete)
}
