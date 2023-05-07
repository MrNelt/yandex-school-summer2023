package order_controller

import (
	log "github.com/bearatol/lg"
	"github.com/labstack/echo/v4"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"yandex-team.ru/bstask/models"
	"yandex-team.ru/bstask/utils"
)

type Storage interface {
	CreateOrder(courier ...models.Order) error
	GetOrderByID(ID int64) (models.Order, error)
	GetOrders(limit, offset int) ([]models.Order, error)
	DeleteOrderByID(ID ...int64) error
	CreateCompleteInfo(info ...models.CompleteInfo) ([]models.Order, error)
	GetOrdersComplete(limit, offset int) ([]models.CompleteInfo, error)
}

type OrderControllerHandler struct {
	storage Storage
}

func NewOrderControllerHandler(s Storage) *OrderControllerHandler {
	return &OrderControllerHandler{storage: s}
}

func (c *OrderControllerHandler) CreateOrders(ctx echo.Context) error {
	rand.Seed(time.Now().UnixNano())
	type query struct {
		Orders []models.Order `json:"orders"`
	}
	orders := new(query)
	if err := ctx.Bind(orders); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	if orders.Orders == nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	for i := range orders.Orders {
		if orders.Orders[i].OrderID == 0 {
			orders.Orders[i].OrderID = rand.Int63n(math.MaxInt64)
		}
	}
	errOperation := c.storage.CreateOrder(orders.Orders...)
	if errOperation != nil {
		log.Error(errOperation)
		return ctx.String(http.StatusBadRequest, errOperation.Error())
	}
	return ctx.JSON(http.StatusOK, orders)
}

func (c *OrderControllerHandler) GetOrderByID(ctx echo.Context) error {
	id, errParse := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if errParse != nil {
		log.Error(errParse)
		return ctx.String(http.StatusBadRequest, errParse.Error())
	}
	order, errOperation := c.storage.GetOrderByID(id)
	if errOperation != nil {
		log.Error(errOperation)
		return ctx.String(http.StatusNotFound, errOperation.Error())
	}
	log.Trace(order)
	return ctx.JSON(http.StatusOK, order)
}

func (c *OrderControllerHandler) GetOrders(ctx echo.Context) error {
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
	orders, errOperation := c.storage.GetOrders(limit, offset)
	if errOperation != nil {
		log.Error(errOperation)
		return ctx.String(http.StatusBadRequest, errOperation.Error())
	}
	log.Trace(orders)
	type answerJson struct {
		Orders []models.Order `json:"orders"`
		Limit  int            `json:"limit"`
		Offset int            `json:"offset"`
	}
	return ctx.JSON(http.StatusOK, &answerJson{orders, limit, offset})
}

func (c *OrderControllerHandler) DeleteOrderByID(ctx echo.Context) error {
	log.Info(ctx.Param("id"))
	id, errParse := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if errParse != nil {
		log.Error(errParse)
		return ctx.String(http.StatusBadRequest, errParse.Error())
	}
	errDel := c.storage.DeleteOrderByID(id)
	if errDel != nil {
		log.Error(errDel)
		return ctx.String(http.StatusBadRequest, errDel.Error())
	}
	return ctx.NoContent(http.StatusOK)
}

func (c *OrderControllerHandler) CompleteOrder(ctx echo.Context) error {
	rand.Seed(time.Now().UnixNano())
	type query struct {
		CompleteInfo []models.CompleteInfo `json:"complete_info"`
	}
	completeInfo := new(query)
	if err := ctx.Bind(completeInfo); err != nil {
		log.Error(err)
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	log.Info(completeInfo.CompleteInfo)
	for _, v := range completeInfo.CompleteInfo {
		if _, err := utils.GetTime(v.CompleteTime); err != nil {
			log.Error(err)
			return ctx.String(http.StatusBadRequest, err.Error())
		}

	}
	orders, errOperation := c.storage.CreateCompleteInfo(completeInfo.CompleteInfo...)
	if errOperation != nil {
		log.Error(errOperation)
		return ctx.String(http.StatusBadRequest, errOperation.Error())
	}
	return ctx.JSON(http.StatusOK, orders)
}

func (c *OrderControllerHandler) GetOrdersComplete(ctx echo.Context) error {
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
	ordersComplete, errOperation := c.storage.GetOrdersComplete(limit, offset)
	if errOperation != nil {
		log.Error(errOperation)
		return ctx.String(http.StatusBadRequest, errOperation.Error())
	}
	log.Trace(ordersComplete)
	type answerJson struct {
		OrdersComplete []models.CompleteInfo `json:"complete_info"`
		Limit          int                   `json:"limit"`
		Offset         int                   `json:"offset"`
	}
	return ctx.JSON(http.StatusOK, &answerJson{ordersComplete, limit, offset})
}
