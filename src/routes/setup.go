package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"yandex-team.ru/bstask/routes/courier_controller"
	"yandex-team.ru/bstask/routes/order_controller"
	"yandex-team.ru/bstask/service"
)

func SetupRoutes(e *echo.Echo) {
	dsn := "postgresql://postgres:password@db/postgres"
	pg, err := service.NewPgService(dsn)
	if err != nil {
		log.Panic(err)
	}

	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))

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
