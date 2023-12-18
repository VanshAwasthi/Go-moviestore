package routes

import (
	"github.com/VanshAwasthi/go-movies/pkg/controller"
	"github.com/gorilla/mux"
)

var RegistermoviesStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/movies/", controller.Createmovies).Methods("POST")
	router.HandleFunc("/movies/", controller.Getmovies).Methods("GET")
	router.HandleFunc("/movies/{moviesId}", controller.GetmoviesById).Methods("GET")
	router.HandleFunc("/movies/{moviesId}", controller.Updatemovies).Methods("PUT")
	router.HandleFunc("/movies/{moviesId}", controller.Deletemovies).Methods("DELETE")
}
