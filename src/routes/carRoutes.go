package routes

import (
	"github.com/gorilla/mux"
	"github.com/mys/go-rentals/src/controllers"
	"github.com/mys/go-rentals/src/utils"

)

var CarsRoutes = func (router *mux.Router){
	router.HandleFunc("/cars", controllers.CreateCar).Methods("POST")
	router.HandleFunc("/addtoken", controllers.AddToken).Methods("POST")
	router.Handle("/testtoken", utils.IsAuthorized(controllers.TestToken) ).Methods("GET")
	router.HandleFunc("/cars", controllers.GetCars).Methods("GET")
	router.HandleFunc("/cars/{registration}/rentals", controllers.RentCar).Methods("POST")
	router.HandleFunc("/cars/{registration}/returns", controllers.ReturnCar).Methods("POST")
}