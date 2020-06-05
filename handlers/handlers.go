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
	router.HandleFunc("/leoTweets", middlewares.ChequeoDB(middlewares.ValidoJWT(routes.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminarTweet", middlewares.ChequeoDB(middlewares.ValidoJWT(routes.EliminarTweet))).Methods("DELETE")

	router.HandleFunc("/subirAvatar", middlewares.ChequeoDB(middlewares.ValidoJWT(routes.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlewares.ChequeoDB(routes.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/subirBanner", middlewares.ChequeoDB(middlewares.ValidoJWT(routes.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenrBanner", middlewares.ChequeoDB(routes.ObtenerBanner)).Methods("GET")

	router.HandleFunc("/altaRelacion", middlewares.ChequeoDB(middlewares.ValidoJWT(routes.AltaRelacion))).Methods("Post")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
