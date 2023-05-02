package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"yandex-team.ru/bstask/models"
)

func createCouriers(ctx echo.Context) error {
	u := new(models.Courier)
	if err := ctx.Bind(u); err != nil {
		return err
	}
	u.ID = "ok!"
	return ctx.JSON(http.StatusOK, u)
}
