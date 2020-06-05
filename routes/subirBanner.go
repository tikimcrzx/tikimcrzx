package routes

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/tikimcrzx/tikimcrzx/bd"
	"github.com/tikimcrzx/tikimcrzx/models"
)

/*SubirBanner subir imagenes de los usuarios de banner al servidor*/
func SubirBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")
	var ex = extension[len(extension)-1]
	var archivo string = "uploads/banners/" + IDUsuario + "." + ex
	print(ex)

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagén ! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagén ! "+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Avatar = IDUsuario + "." + ex
	status, err = bd.ModificoRegistro(usuario, IDUsuario)
	if err != nil || status == false {
		http.Error(w, "Error al grabar el banner ! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
