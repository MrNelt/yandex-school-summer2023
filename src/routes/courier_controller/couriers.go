package courier_controller

import (
	log "github.com/bearatol/lg"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"yandex-team.ru/bstask/models"
)

type Storage interface {
	CreateCourier(courier ...models.Courier) error
	GetCourierByID(ID int64) (models.Courier, error)
	GetCouriers(limit, offset int) ([]models.Courier, error)
	DeleteCourierByID(ID ...int64) error
}

type CourierControllerHandler struct {
	storage Storage
}

func NewCourierControllerHandler(s Storage) *CourierControllerHandler {
	return &CourierControllerHandler{storage: s}
}

func (c *CourierControllerHandler) CreateCouriers(ctx echo.Context) error {
	type query struct {
		Couriers []models.Courier `json:"couriers"`
	}
	couriers := new(query)
	if err := ctx.Bind(couriers); err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	errOperation := c.storage.CreateCourier(couriers.Couriers...)
	if errOperation != nil {
		log.Error(errOperation)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	return ctx.String(http.StatusOK, "ok")
}

func (c *CourierControllerHandler) GetCourierByID(ctx echo.Context) error {
	id, errParse := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if errParse != nil {
		log.Error(errParse)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	courier, errOperation := c.storage.GetCourierByID(id)
	if errOperation != nil {
		log.Error(errOperation)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	log.Trace(courier)
	return ctx.JSON(http.StatusOK, courier)
}

func (c *CourierControllerHandler) GetCouriers(ctx echo.Context) error {
	limitStr := ctx.QueryParam("limit")
	offsetStr := ctx.QueryParam("offset")
	if limitStr == "" {
		limitStr = "1"
	}
	if offsetStr == "" {
		offsetStr = "0"
	}
	limit, errLimit := strconv.Atoi(limitStr)
	offset, errOffset := strconv.Atoi(offsetStr)
	log.Infof("%d %d", limit, offset)
	if errLimit != nil {
		log.Error(errLimit)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	if errOffset != nil {
		log.Error(errLimit)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	couriers, errOperation := c.storage.GetCouriers(limit, offset)
	if errOperation != nil {
		log.Error(errOperation)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	log.Trace(couriers)
	type answerJson struct {
		Couriers []models.Courier `json:"couriers"`
		Limit    int              `json:"limit"`
		Offset   int              `json:"offset"`
	}
	return ctx.JSON(http.StatusOK, &answerJson{couriers, limit, offset})
}

func (c *CourierControllerHandler) DeleteCourierByID(ctx echo.Context) error {
	log.Info(ctx.Param("id"))
	id, errParse := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if errParse != nil {
		log.Error(errParse)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	errDel := c.storage.DeleteCourierByID(id)
	if errDel != nil {
		log.Error(errDel)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	return ctx.JSON(http.StatusOK, "ok")
}
