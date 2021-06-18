package routers

import (
	"errors"
	"strings"

	"github.com/aledaas/tiwtgo/src/github.com/aledaas/tiwtgo/bd"
	"github.com/aledaas/tiwtgo/src/github.com/aledaas/tiwtgo/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Email logueado*/
var Email string

/*IDUsuario logueado*/
var IDUsuario string

/*ProcesoToken procesa el token enviado*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("mi_semilla")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de tocken invalido")
	}
	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, ID := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = ID
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, string(""), err
}
