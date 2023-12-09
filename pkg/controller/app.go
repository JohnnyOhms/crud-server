package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/JohnnyOhms/crud-server/pkg/config"
	"github.com/JohnnyOhms/crud-server/pkg/model"
	"github.com/JohnnyOhms/crud-server/pkg/utils"
	_ "github.com/gorilla/mux"
)

type Result struct {
	Success bool
	Mssg    string
	Data    interface{}
}

func CreateInfo(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil || r.ContentLength == 0 {
		http.Error(w, "Empty request body", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	collection, err := config.ConnectDB()
	if err != nil {
		http.Error(w, "failed to establish database connection", http.StatusInternalServerError)
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer collection.Database().Client().Disconnect(context.Background())

	var info model.Info
	err = utils.ParseBody(r, &info)
	if err != nil {
		http.Error(w, "failed to parse body", http.StatusBadRequest)
		fmt.Println("Error parsing request body:", err)
		return
	}

	res, err := model.InsertInfo(info, collection)
	if err != nil {
		http.Error(w, "failed to add to database", http.StatusInternalServerError)
		fmt.Println("Error inserting data into database:", err)
		return
	}
	data := Result{
		Success: true,
		Mssg:    "added successfully",
		Data:    res,
	}
	response, err := utils.Encoding(data)
	if err != nil {
		http.Error(w, "failed to encode data", http.StatusInternalServerError)
		fmt.Println("Error encoding response:", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func GetData(w http.ResponseWriter, r *http.Request) {
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
