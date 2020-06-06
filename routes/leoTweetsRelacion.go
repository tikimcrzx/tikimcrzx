package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/tikimcrzx/tikimcrzx/bd"
)

/*LeoTweetsSeguidores lee los tweets de los deguidores del usuario*/
func LeoTweetsSeguidores(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}
	temPagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página como entero mayor a 0 "+err.Error(), http.StatusBadRequest)
		return
	}

	pagina := int64(temPagina)

	respuesta, correcto := bd.LeoTweetsSeguidores(IDUsuario, pagina)
	if correcto == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respuesta)
}
