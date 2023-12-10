package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/JohnnyOhms/crud-server/pkg/config"
	"github.com/JohnnyOhms/crud-server/pkg/model"
	"github.com/JohnnyOhms/crud-server/pkg/utils"
	"github.com/gorilla/mux"
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

func GetInfo(w http.ResponseWriter, r *http.Request) {
	var userId string
	if userId == "" || len(userId) == 0 {
		http.Error(w, "authenticate to continue ", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type:", "aplicaton/json")

	res, err := model.GetInfo(userId, config.Collection)
	if err != nil {
		http.Error(w, "failed to get data from database", http.StatusBadRequest)
		fmt.Println("Error in getting data from db", err)
		return
	}
	data := Result{
		Success: true,
		Mssg:    "info gotten successfully",
		Data:    res,
	}

	result, err := utils.Encoding(data)
	if err != nil {
		http.Error(w, "failed to encode data", http.StatusInternalServerError)
		fmt.Println("Error in encoding data", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetSingleInfo(w http.ResponseWriter, r *http.Request) {
	var userId string
	if userId == "" || len(userId) == 0 {
		http.Error(w, "authenticate to continue ", http.StatusUnauthorized)
		return
	}
	params := mux.Vars(r)
	if params["id"] == "" {
		http.Error(w, "id param cannot be null", http.StatusBadRequest)
		fmt.Println("null id param provided")
		return
	}
	Id := params["id"]
	res, err := model.GetSingleInfo(userId, Id, config.Collection)

	if err != nil {
		http.Error(w, "failed to get data from database", http.StatusBadRequest)
		fmt.Println("Error in getting data from db", err)
		return
	}
	data := Result{
		Success: true,
		Mssg:    "info gotten successfully",
		Data:    res,
	}

	result, err := utils.Encoding(data)
	if err != nil {
		http.Error(w, "failed to encode data", http.StatusInternalServerError)
		fmt.Println("Error in encoding data", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func EditInfo(w http.ResponseWriter, r *http.Request) {
	var userId string
	if userId == "" || len(userId) == 0 {
		http.Error(w, "authenticate to continue ", http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	if params["id"] == "" {
		http.Error(w, "id param cannot be null", http.StatusBadRequest)
		fmt.Println("null id param provided")
		return
	}
	Id := params["id"]

	if r.Body == nil || r.ContentLength == 0 {
		http.Error(w, "Empty request body", http.StatusBadRequest)
		return
	}
	var update interface{}

	updated := utils.ParseBody(r, &update)

	result, err := model.EditInfo(userId, Id, updated, config.Collection)
	if err != nil {
		http.Error(w, "failed to update data from the data base", http.StatusBadRequest)
		return
	}

	data := Result{
		Success: true,
		Mssg:    "info updated successfully",
		Data:    result,
	}

	res, err := utils.Encoding(data)
	if err != nil {
		http.Error(w, "failed to encode data ", http.StatusInternalServerError)
		fmt.Println("encoding data failed")
		return
	}
	w.Header().Set("Content-Type:", "aplicaton/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)

}

func DeleteSingleInfo(w http.ResponseWriter, r *http.Request) {
	var userId string
	if userId == "" || len(userId) == 0 {
		http.Error(w, "authenticate to continue ", http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	if params["id"] == "" {
		http.Error(w, "id param cannot be null", http.StatusBadRequest)
		fmt.Println("null id param provided")
		return
	}
	Id := params["id"]

	result, err := model.DeleteSingleInfo(userId, Id, config.Collection)
	if err != nil {
		http.Error(w, "failed to update data from the data base", http.StatusBadRequest)
		return
	}

	data := Result{
		Success: true,
		Mssg:    "info updated successfully",
		Data:    result,
	}

	res, err := utils.Encoding(data)
	if err != nil {
		http.Error(w, "failed to encode data ", http.StatusInternalServerError)
		fmt.Println("encoding data failed")
		return
	}

	w.Header().Set("Content-Type:", "aplicaton/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)

}

func DeleteInfo(w http.ResponseWriter, r *http.Request) {
	var userId string
	if userId == "" || len(userId) == 0 {
		http.Error(w, "authenticate to continue ", http.StatusUnauthorized)
		return
	}
	res, err := model.DeleteInfo(userId, config.Collection)
	if err != nil {
		http.Error(w, "failed to delete info", http.StatusInternalServerError)
		return
	}

	data := Result{
		Success: true,
		Mssg:    "all infos deleted sucessfully",
		Data:    res,
	}

	result, err := utils.Encoding(data)
	if err != nil {
		http.Error(w, "failed to encode info", http.StatusInternalServerError)
		fmt.Println("encoding data failed")
		return
	}

	w.Header().Set("Content-Type:", "aplicaton/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(result)
}
