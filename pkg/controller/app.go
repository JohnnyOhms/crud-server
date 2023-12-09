package controller

import (
	"net/http"

	_ "github.com/gorilla/mux"
)

func GetData( w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type:", "aplicaton/json")
	w.WriteHeader(http.StatusOK)

}

func GetSingleData(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// Id := params["id"]
	w.Header().Set("Content-Type:", "aplicaton/json")
	w.WriteHeader(http.StatusOK)

}

func EditData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type:", "aplicaton/json")
	w.WriteHeader(http.StatusCreated)

}

func DeleteSingleData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type:", "aplicaton/json")
	w.WriteHeader(http.StatusCreated)

}

func DeleteData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type:", "aplicaton/json")
	w.WriteHeader(http.StatusCreated)

}
