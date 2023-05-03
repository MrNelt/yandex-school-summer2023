package storage

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"os"
	"yandex-team.ru/bstask/models"
)

type Storage interface {
	createCourier(courier ...models.Courier) error
	getCourierByID(ID int64) (models.Courier, error)
	getCouriers(limit, offset int) ([]models.Courier, error)
	deleteCourierByID(ID ...int64) error
}

var storage Storage

func Init() {
	dsl := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_SSLMODE"))
	log.Print(dsl)
	var err error
	storage, err = NewPgStorage(dsl)
	if err != nil {
		log.Panic(err)
	}
}

func CreateCourier(courier ...models.Courier) error {
	return storage.createCourier(courier...)
}

func GetCourierByID(ID int64) (models.Courier, error) {
	return storage.getCourierByID(ID)
}

func GetCouriers(limit, offset int) ([]models.Courier, error) {
	return storage.getCouriers(limit, offset)
}
