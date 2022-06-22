package router

import (
	"github.com/gorilla/mux"
	"github.com/niladridas/mongoAPI/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", controller.ServeHome).Methods("GET")
	router.HandleFunc("api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("api/movie/{id}", controller.MarkedAsWatched).Methods("PUT")
	router.HandleFunc("api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("api/deleteallmovie", controller.DeleteAllMovie).Methods("DELETE")

	return router
}
