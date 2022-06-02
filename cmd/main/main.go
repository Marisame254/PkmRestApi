package main

import (
	"log"
	"net/http"

	"github.com/Marisame254/PkmRestApi/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterPkmRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:2540", r))
}
