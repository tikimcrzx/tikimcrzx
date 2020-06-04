package middlewares

import (
	"net/http"

	"github.com/tikimcrzx/tikimcrzx/bd"
)

/*ChequeoDB middleware verifica conexion*/
func ChequeoDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConexion() == 0 {
			http.Error(w, "Conexión perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
