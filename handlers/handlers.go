package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/tikimcrzx/tikimcrzx/middlewares"
	"github.com/tikimcrzx/tikimcrzx/routes"
)

/*Manejadores seteo puerto y mando a escuchar servidor*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlewares.ChequeoDB(routes.Registro)).Methods("POST")
	router.HandleFunc("/login", middlewares.ChequeoDB(routes.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlewares.ChequeoDB(middlewares.ValidoJWT(routes.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlewares.ChequeoDB(middlewares.ValidoJWT(routes.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlewares.ChequeoDB(middlewares.ValidoJWT(routes.GraboTweet))).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
