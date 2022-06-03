package main

import (
	"log"
	"net/http"

	"github.com/Marisame254/PkmRestApi/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterPkmRoutes(router)
	http.Handle("/", router)
	log.Println("\nAPI is running on port 4000")
	http.ListenAndServe("localhost:4000", router)
}
