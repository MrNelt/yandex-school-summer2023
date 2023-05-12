package service

import (
	"encoding/json"
	"fmt"
	log "github.com/bearatol/lg"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
	"yandex-team.ru/bstask/models"
	"yandex-team.ru/bstask/utils"
)

type PgService struct {
	db *gorm.DB
}

func NewPgService(dsn string) (*PgService, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	for err != nil {
		log.Info("Trying to connect to pg")
		time.Sleep(2 * time.Second)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}
	log.Info("Success connecting to pg")
	// restrictions on the task
	err = db.AutoMigrate(&models.CourierDB{}, &models.OrderDB{}, &models.CompleteInfo{})
	if err != nil {
		return nil, err
	}
	return &PgService{db}, nil
}

func (p *PgService) CreateCourier(courier ...models.Courier) error {
	db := p.db
	tx := db.Begin()
	for _, v := range courier {
		courierJson, errJs := json.Marshal(v)
		if errJs != nil {
			return errJs
		}
		errCreate := tx.Create(&models.CourierDB{ID: v.CourierID, Attributes: courierJson}).Error
		if errCreate != nil {
			tx.Rollback()
			return errCreate
		}
	}
	tx.Commit()
	return nil
}

func (p *PgService) GetCourierByID(ID int64) (models.Courier, error) {
	db := p.db
	courierDB := models.CourierDB{}
	if result := db.Limit(1).Find(&courierDB, ID); result.RowsAffected == 0 {
		return models.Courier{}, fmt.Errorf("courier doesn't exist")
	}
	courier := models.Courier{}
	if err := json.Unmarshal(courierDB.Attributes, &courier); err != nil {
		return models.Courier{}, err
	}
	return courier, nil
}

func (p *PgService) GetCouriers(limit, offset int) ([]models.Courier, error) {
	db := p.db
	var couriersDB []models.CourierDB
	var couriers []models.Courier
	db.Limit(limit).Offset(offset).Find(&couriersDB)
	for _, courierDB := range couriersDB {
		courier := models.Courier{}
		err := json.Unmarshal(courierDB.Attributes, &courier)
		if err != nil {
			return nil, err
		}
		couriers = append(couriers, courier)
	}
	return couriers, nil
}

func (p *PgService) DeleteCourierByID(ID ...int64) error {
	db := p.db
	tx := db.Begin()
	for _, id := range ID {
		err := tx.Delete(models.CourierDB{}, id).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func (p *PgService) CreateOrder(order ...models.Order) error {
	db := p.db
	tx := db.Begin()
	for _, v := range order {
		orderJson, errJs := json.Marshal(v)
		if errJs != nil {
			return errJs
		}
		errCreate := tx.Create(&models.OrderDB{ID: v.OrderID, Attributes: orderJson}).Error
		if errCreate != nil {
			tx.Rollback()
			return errCreate
		}
	}
	tx.Commit()
	return nil
}

func (p *PgService) GetOrderByID(ID int64) (models.Order, error) {
	db := p.db
	orderDB := models.OrderDB{}
	if result := db.Limit(1).Find(&orderDB, ID); result.RowsAffected == 0 {
		return models.Order{}, fmt.Errorf("order doesn't exist")
	}
	order := models.Order{}
	if err := json.Unmarshal(orderDB.Attributes, &order); err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func (p *PgService) GetOrders(limit, offset int) ([]models.Order, error) {
	db := p.db
	var ordersDB []models.OrderDB
	var orders []models.Order
	if result := db.Limit(limit).Offset(offset).Find(&ordersDB); result.Error != nil {
		return nil, result.Error
	}
	for _, orderDB := range ordersDB {
		order := models.Order{}
		err := json.Unmarshal(orderDB.Attributes, &order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (p *PgService) DeleteOrderByID(ID ...int64) error {
	db := p.db
	tx := db.Begin()
	for _, id := range ID {
		err := tx.Delete(models.OrderDB{}, id).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func (p *PgService) CreateCompleteInfo(info ...models.CompleteInfo) ([]models.Order, error) {
	db := p.db
	tx := db.Begin()
	var orders []models.Order
	for _, v := range info {
		log.Info(v)
		completeInfo, errExistCompleteInfo := p.GetOrderCompleteByID(v.OrderId)
		if errExistCompleteInfo == nil {
			if completeInfo == v {
				order, _ := p.GetOrderByID(v.OrderId)
				orders = append(orders, order)
				continue
			} else {
				return nil, fmt.Errorf("bad request")
			}
		}
		_, errExistCourier := p.GetCourierByID(v.CourierId)
		if errExistCourier != nil {
			tx.Rollback()
			return nil, fmt.Errorf("don't exist courier")
		}
		order, errExistOrder := p.GetOrderByID(v.OrderId)
		if errExistOrder != nil {
			tx.Rollback()
			return nil, fmt.Errorf("dont't exist order")
		}
		if result := tx.Create(&v); result.Error != nil {
			tx.Rollback()
			return nil, result.Error
		}
		order.CompletedTime = v.CompleteTime
		orderJson, errJs := json.Marshal(order)
		if errJs != nil {
			tx.Rollback()
			return nil, errJs
		}
		if result := tx.Model(&models.OrderDB{ID: v.OrderId}).Update("attributes", orderJson); result.Error != nil {
			tx.Rollback()
			return nil, result.Error
		}
		orders = append(orders, order)
	}
	tx.Commit()
	return orders, nil
}

func (p *PgService) GetOrdersComplete(limit, offset int) ([]models.CompleteInfo, error) {
	db := p.db
	var info []models.CompleteInfo
	if result := db.Limit(limit).Offset(offset).Find(&info); result.Error != nil {
		return nil, result.Error
	}
	return info, nil
}

func (p *PgService) GetOrderCompleteByID(ID int64) (models.CompleteInfo, error) {
	db := p.db
	completeInfo := models.CompleteInfo{}
	if result := db.First(&completeInfo, ID); result.Error != nil {
		return models.CompleteInfo{}, result.Error
	}
	return completeInfo, nil
}

func (p *PgService) CalculateRatingCourier(ID int64, startDate, endDate time.Time) (models.Courier, int32, int32, error) {
	db := p.db
	courier, err := p.GetCourierByID(ID)
	var cfEarning int32 = 0
	var cfRating int32 = 0
	switch courier.CourierType {
	case "FOOT":
		cfEarning = 2
		cfRating = 3
	case "BIKE":
		cfEarning = 3
		cfRating = 2
	case "AUTO":
		cfEarning = 4
		cfRating = 1
	}

	if err != nil {
		return models.Courier{}, 0, 0, err
	}
	var info []models.CompleteInfo
	if result := db.Where("courier_id = ?", ID).Find(&info); result.Error != nil {
		return models.Courier{}, 0, 0, result.Error
	}
	var countCompleteOrders int32 = 0
	var earning int32 = 0
	for _, i := range info {
		date, errParse := utils.GetTime(i.CompleteTime)
		if err != nil {
			return models.Courier{}, 0, 0, errParse
		}
		if !utils.InsideInterval(date, startDate, endDate) {
			continue
		}
		countCompleteOrders++
		order, errOperation := p.GetOrderByID(i.OrderId)
		if errOperation != nil {
			return models.Courier{}, 0, 0, errParse
		}
		earning += order.Cost
	}
	earning *= cfEarning
	var hours = int32((endDate.Unix() - startDate.Unix()) / 3600)
	if hours == 0 {
		hours = 1
	}
	var rating = (countCompleteOrders / hours) * cfRating
	return courier, rating, earning, nil
}
