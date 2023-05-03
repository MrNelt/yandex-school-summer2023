package courier_conroller

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
	"yandex-team.ru/bstask/models"
	"yandex-team.ru/bstask/storage"
)

func CreateCouriers(ctx echo.Context) error {
	type query struct {
		Couriers []models.Courier `json:"couriers"`
	}
	couriers := new(query)
	if err := ctx.Bind(couriers); err != nil {
		return err
	}
	log.Info(couriers)
	err := storage.CreateCourier(couriers.Couriers...)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	return ctx.String(http.StatusOK, "ok")
}

func GetCourierByID(ctx echo.Context) error {
	id, errParse := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if errParse != nil {
		return errParse
	}
	courier, errOperation := storage.GetCourierByID(id)
	if errOperation != nil {
		return errOperation
	}
	return ctx.JSON(http.StatusOK, courier)
}

func GetCouriers(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "GetCouriers")
}
func DeleteCourierByID(ctx echo.Context) error {
	log.Info(ctx.Param("id"))
	return ctx.JSON(http.StatusOK, "DeleteCouriers")
}
