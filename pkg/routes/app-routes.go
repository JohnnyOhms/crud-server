package routes

import (
	"github.com/JohnnyOhms/crud-server/pkg/controller"
	"github.com/gorilla/mux"
)

func Router(routes *mux.Router) {
	routes.HandleFunc("/getdata", controller.GetData).Methods("GET")
	routes.HandleFunc("/getsingledata/{id}", controller.GetSingleData).Methods("GET")
	routes.HandleFunc("/editdata/{id}", controller.EditData).Methods("PUT")
	routes.HandleFunc("/deletesingledata/{id}", controller.DeleteSingleData).Methods("DELETE")
	routes.HandleFunc("/deletedata", controller.DeleteData).Methods("DELETE")
}
