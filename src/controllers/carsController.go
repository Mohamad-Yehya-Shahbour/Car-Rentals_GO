package controllers

import (
	"encoding/json"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/mys/go-rentals/src/models"
	"github.com/mys/go-rentals/src/utils"
)

var NewCar models.Car

type Temp struct {
	Driven int `json:"driven"`
}


func GetCars(w http.ResponseWriter, r *http.Request) {
	NewCar := models.GetAllCars()
	res, _ := json.Marshal(NewCar)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func CreateCar(w http.ResponseWriter, r *http.Request) {
	CreateCar := &models.Car{}
	utils.ParseBody(r, CreateCar)
	CreateCar.Status = "available"
	FoundedCar, _ := models.GetCarByRg(CreateCar.Registration)

	if FoundedCar.Registration != "" {
		http.Error(w, "the registration number already exists", http.StatusBadRequest)
		return
	}

	c := CreateCar.CreateCar()
	var id string = strconv.FormatUint(uint64(c.ID), 10)
	utils.ResponseHandler(w, "CarId", id, 200)
}


func RentCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reg := vars["registration"]
	FoundedCar, db := models.GetCarByRg(reg)

	if FoundedCar.Registration == "" {
		http.Error(w, "this car does not exists", http.StatusBadRequest)
		return
	}

	if FoundedCar.Status == "rented" {
		http.Error(w, "this car is already rented", http.StatusBadRequest)
		return
	}

	FoundedCar.Status = "rented"
	db.Save(&FoundedCar)
	utils.ResponseHandler(w, "message", "status updated to rented", 200)
}


func ReturnCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reg := vars["registration"]
	FoundedCar, db := models.GetCarByRg(reg)

	if FoundedCar.Registration == "" {
		http.Error(w, "this car does not exists", http.StatusBadRequest)
		return
	}

	if FoundedCar.Status == "available" {
		http.Error(w, "this car is not rented, there is a problem! please report this issue ", http.StatusBadRequest)
		return
	}

	var temp Temp

	// driven kilometers should be entered this way as json data through the request

	// {
	// 	"driven":7000
	// }
	
	err := json.NewDecoder(r.Body).Decode(&temp)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	FoundedCar.Mileage = FoundedCar.Mileage + temp.Driven
	FoundedCar.Status = "available"
	db.Save(&FoundedCar)
	utils.ResponseHandler(w, "message", "car status updated to available, mileage updated", 200)
}


