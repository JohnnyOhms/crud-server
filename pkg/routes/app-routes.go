package routes

import (
	"github.com/JohnnyOhms/crud-server/pkg/controller"
	"github.com/gorilla/mux"
)

func Router(routes *mux.Router) {
	routes.HandleFunc("/createinfo", controller.CreateInfo).Methods("POST")
	routes.HandleFunc("/getinfo", controller.GetData).Methods("GET")
	routes.HandleFunc("/getsingleinfo/{id}", controller.GetSingleData).Methods("GET")
	routes.HandleFunc("/editinfo/{id}", controller.EditData).Methods("PUT")
	routes.HandleFunc("/deletesingleinfo/{id}", controller.DeleteSingleData).Methods("DELETE")
	routes.HandleFunc("/deleteinfo", controller.DeleteData).Methods("DELETE")
}
