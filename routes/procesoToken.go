package routes

import (
	"errors"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/tikimcrzx/tikimcrzx/bd"
	"github.com/tikimcrzx/tikimcrzx/models"
)

/*Email email en el token*/
var Email string

/*IDUsuario id del usuario*/
var IDUsuario string

/*ProcesoToken procesa token del usuario*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	secret := os.Getenv("jwt_secret")
	if secret == "" {
		secret = "My_Secrect_Key_Chingona"
	}
	miClave := []byte(secret)
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err
}
