package jwt

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/tikimcrzx/tikimcrzx/models"
)

/*GeneroJWT genera jwt con datos necesarios*/
func GeneroJWT(t models.Usuario) (string, error) {
	secret := os.Getenv("jwt_secret")
	if secret == "" {
		secret = "My_Secrect_Key_Chingona"
	}

	miClave := []byte(secret)
	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
