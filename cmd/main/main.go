package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/JohnnyOhms/crud-server/pkg/Doc"
	Config "github.com/JohnnyOhms/crud-server/pkg/config"
	"github.com/JohnnyOhms/crud-server/pkg/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.Router(r)
	r.HandleFunc("/", Doc.Documentation).Methods("GET")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	corsRouter := handlers.CORS(headersOk, originsOk, methodsOk)(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if _, err := Config.ConnectDB(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Server listening on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, corsRouter); err != nil {
		log.Fatal(err)
	}
}
