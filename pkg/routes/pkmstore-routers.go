package routes

import (
	"github.com/Marisame254/PkmRestApi/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterPkmRoutes = func(router *mux.Router) {
	router.HandleFunc("/pokemon/", controllers.CreatePkm).Methods("POST")
	router.HandleFunc("/pokemon/", controllers.GetPkm).Methods("GET")
	router.HandleFunc("/pokemon/{Id}", controllers.GetPkmById).Methods("GET")
	router.HandleFunc("/pokemon/{Id}", controllers.UpdatePkm).Methods("PUT")
	router.HandleFunc("/pokemon/{Id}", controllers.DeletePkm).Methods("DELETE")
}
