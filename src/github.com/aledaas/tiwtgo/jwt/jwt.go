package jwt

import (
	"time"

	"github.com/aledaas/tiwtgo/src/github.com/aledaas/tiwtgo/models"

	jwt "github.com/dgrijalva/jwt-go"
)

/*GeneroJWT genera el jwt*/
func GeneroJWT(t models.Usuario) (string, error) {
	miClave := []byte("mi_semilla")
	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}
	tocken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tockenStr, err := tocken.SignedString(miClave)
	if err != nil {
		return tockenStr, err
	}
	return tockenStr, nil
}
