package routes

import (
	"github.com/JohnnyOhms/crud-server/pkg/controller"
	"github.com/gorilla/mux"
)

func Router(routes *mux.Router) {
	routes.HandleFunc("/createinfo", controller.CreateInfo).Methods("POST")
	routes.HandleFunc("/getinfo/{userid}", controller.GetInfo).Methods("GET")
	routes.HandleFunc("/getsingleinfo/{userid}/{id}", controller.GetSingleInfo).Methods("GET")
	routes.HandleFunc("/editinfo/{userid}/{id}", controller.EditInfo).Methods("PUT")
	routes.HandleFunc("/deletesingleinfo/{userid}/{id}", controller.DeleteSingleInfo).Methods("DELETE")
	routes.HandleFunc("/deleteinfo/{userid}", controller.DeleteInfo).Methods("DELETE")
}
