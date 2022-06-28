package routes

import (
	"github.com/gorilla/mux"
	"github.com/mys/go-rentals/src/controllers"
)

var CarsRoutes = func (router *mux.Router){
	router.HandleFunc("/cars", controllers.CreateCar).Methods("POST")
	router.HandleFunc("/cars", controllers.GetCars).Methods("GET")
	router.HandleFunc("/cars/{registration}/rentals", controllers.RentCar).Methods("POST")
	router.HandleFunc("/cars/{registration}/returns", controllers.ReturnCar).Methods("POST")
}