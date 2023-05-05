package courier_controller

import (
	log "github.com/bearatol/lg"
	"github.com/labstack/echo/v4"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
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

// @Summary create couriers
// @Tags courier-controller
// @Description create couriers
// @Accept  json
// @Produce  json
// @Param input body query true "couriers info"
// @Success 200 {object} query
// @Failure 400 {string] string
// @Router /couriers [post]

func (c *CourierControllerHandler) CreateCouriers(ctx echo.Context) error {
	rand.Seed(time.Now().UnixNano())
	type query struct {
		Couriers []models.Courier `json:"couriers"`
	}
	couriers := new(query)
	if err := ctx.Bind(couriers); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	if couriers.Couriers == nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	for i := range couriers.Couriers {
		if couriers.Couriers[i].CourierID == 0 {
			couriers.Couriers[i].CourierID = rand.Int63n(math.MaxInt64)
		}
	}
	errOperation := c.storage.CreateCourier(couriers.Couriers...)
	if errOperation != nil {
		log.Error(errOperation)
		return ctx.String(http.StatusBadRequest, errOperation.Error())
	}
	return ctx.JSON(http.StatusOK, couriers)
}

func (c *CourierControllerHandler) GetCourierByID(ctx echo.Context) error {
	id, errParse := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if errParse != nil {
		log.Error(errParse)
		return ctx.String(http.StatusBadRequest, errParse.Error())
	}
	courier, errOperation := c.storage.GetCourierByID(id)
	if errOperation != nil {
		log.Error(errOperation)
		return ctx.String(http.StatusNotFound, errOperation.Error())
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
		return ctx.String(http.StatusBadRequest, errLimit.Error())
	}
	if errOffset != nil {
		log.Error(errOffset)
		return ctx.String(http.StatusBadRequest, errOffset.Error())
	}
	couriers, errOperation := c.storage.GetCouriers(limit, offset)
	if errOperation != nil {
		log.Error(errOperation)
		return ctx.String(http.StatusBadRequest, errOperation.Error())
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
		return ctx.String(http.StatusBadRequest, errParse.Error())
	}
	errDel := c.storage.DeleteCourierByID(id)
	if errDel != nil {
		log.Error(errDel)
		return ctx.String(http.StatusBadRequest, errDel.Error())
	}
	return ctx.NoContent(http.StatusOK)
}
