package main

import (
	"fmt"
	"log"
	"net/http"

	Config "github.com/JohnnyOhms/crud-server/pkg/config"
	"github.com/JohnnyOhms/crud-server/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.Router(r)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type:", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "<h1 style='color:red'>DOCUMENTATION</h1>")
	})

	if _, err := Config.ConnectDB(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("server running on port 8080\n")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
