package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ping(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "pong")
}
