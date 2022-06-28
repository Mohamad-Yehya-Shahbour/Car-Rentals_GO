package models

import (
	"github.com/mys/go-rentals/src/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Car struct {
	gorm.Model
	CarModel       	string `json:"model"`
	Registration   	string `gorm:"" json:"registration"`
	Mileage			int `json:"mileage"`
	Status			string `json:"status"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Car{})
}


func (c *Car) CreateCar() *Car {
	//db.NewRecord(b)
	db.Create(&c)
	return c
}

func GetAllCars() []Car {
	var cars []Car
	db.Find(&cars)
	return cars
}

func GetCarById(Id int64) (*Car, *gorm.DB) {
	var getCar Car
	db := db.Where("ID=?", Id).Find(&getCar)
	return &getCar, db
}

func GetCarByRg(rg string) (*Car, *gorm.DB) {
	var getCar Car
	db := db.Where("registration=?", rg).Find(&getCar)
	return &getCar, db
}

func DeleteCar(Id int64) Car {
	var car Car
	db.Where("ID=?", Id).Delete(&car)
	return car
}