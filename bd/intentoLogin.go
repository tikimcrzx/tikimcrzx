package bd

import (
	"github.com/tikimcrzx/tikimcrzx/models"
	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin realiza verificación de credenciales*/
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usuario, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return usuario, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(usuario.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)

	if err != nil {
		return usuario, false
	}

	return usuario, true
}
