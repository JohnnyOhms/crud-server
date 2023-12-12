package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/JohnnyOhms/crud-server/pkg/config"
	"github.com/JohnnyOhms/crud-server/pkg/model"
	"github.com/JohnnyOhms/crud-server/pkg/utils"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil || r.ContentLength == 0 {
		http.Error(w, "Empty request body", http.StatusBadRequest)
		return
	}

	collection, err := config.ConnectAuth()
	if err != nil {
		http.Error(w, "failed to establish database connection", http.StatusInternalServerError)
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer collection.Database().Client().Disconnect(context.Background())

	var user model.User
	err = utils.ParseBody(r, &user)
	if err != nil {
		http.Error(w, "failed to parse body", http.StatusBadRequest)
		fmt.Println("Error parsing request body:", err)
		return
	}
	user.UserId = utils.GenerateUserId()
	err = model.CreateUser(user, collection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error creating user:", err)
		return
	}

	response, err := utils.Encoding(user)
	if err != nil {
		http.Error(w, "failed to encode data", http.StatusInternalServerError)
		fmt.Println("Error encoding response:", err)
		return
	}

	w.Header().Set("Content-Type:", "applicaton/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func Signin(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil || r.ContentLength == 0 {
		http.Error(w, "Empty request body", http.StatusBadRequest)
		return
	}

	collection, err := config.ConnectAuth()
	if err != nil {
		http.Error(w, "failed to establish database connection", http.StatusInternalServerError)
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer collection.Database().Client().Disconnect(context.Background())

	var user model.LoginUser
	err = utils.ParseBody(r, &user)
	if err != nil {
		http.Error(w, "failed to parse body", http.StatusBadRequest)
		fmt.Println("Error parsing request body:", err)
		return
	}
	res, err := model.GetUser(user, collection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error creating user:", err)
		return
	}
	response, err := utils.Encoding(res)
	if err != nil {
		http.Error(w, "failed to encode data", http.StatusInternalServerError)
		fmt.Println("Error encoding response:", err)
		return
	}

	w.Header().Set("Content-Type:", "applicaton/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
