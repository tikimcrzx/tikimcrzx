package routes

import (
	"encoding/json"
	"net/http"

	"github.com/tikimcrzx/tikimcrzx/bd"
	"github.com/tikimcrzx/tikimcrzx/models"
)

/*ModificarPerfil ruta para modificar perfil*/
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), http.StatusBadRequest)
		return
	}

	var status bool

	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar modificar el registro "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado modificar el registro", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
