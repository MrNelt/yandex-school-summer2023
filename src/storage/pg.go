package storage

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
	"yandex-team.ru/bstask/models"
)

type PgStorage struct {
	db *gorm.DB
}

func NewPgStorage(dsl string) (*PgStorage, error) {
	db, err := gorm.Open(postgres.Open(dsl), &gorm.Config{})
	for err != nil {
		log.Info("Trying to connect to pg")
		time.Sleep(2 * time.Second)
		db, err = gorm.Open(postgres.Open(dsl), &gorm.Config{})
	}
	log.Info("Success connecting to pg")
	err = db.AutoMigrate(&models.CourierDB{})
	if err != nil {
		return nil, err
	}
	return &PgStorage{db}, nil
}

func (p *PgStorage) CreateCourier(courier ...models.Courier) error {
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

func (p *PgStorage) GetCourierByID(ID int64) (models.Courier, error) {
	db := p.db
	courierDB := models.CourierDB{}
	if result := db.First(&courierDB, ID); result.Error != nil {
		return models.Courier{}, result.Error
	}
	courier := models.Courier{}
	if err := json.Unmarshal(courierDB.Attributes, &courier); err != nil {
		return models.Courier{}, err
	}
	return courier, nil
}

func (p *PgStorage) GetCouriers(limit, offset int) ([]models.Courier, error) {
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

func (p *PgStorage) DeleteCourierByID(ID ...int64) error {
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
