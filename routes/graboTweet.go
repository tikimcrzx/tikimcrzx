package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/tikimcrzx/tikimcrzx/bd"
	"github.com/tikimcrzx/tikimcrzx/models"
)

/*GraboTweet*/
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrión un error al intentar insertar el registro, reintente nuevamente "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertal el tweet", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
