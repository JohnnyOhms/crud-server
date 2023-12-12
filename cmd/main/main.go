package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JohnnyOhms/crud-server/pkg/Doc"
	Config "github.com/JohnnyOhms/crud-server/pkg/config"
	"github.com/JohnnyOhms/crud-server/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.Router(r)
	r.HandleFunc("/", Doc.Documentation).Methods("GET")

	if _, err := Config.ConnectDB(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("server running on port 8080\n")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
