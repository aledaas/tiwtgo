package middlew

import (
	"net/http"

	"github.com/aledaas/tiwtgo/bd"
)

/*ChequeoBD es el middleware que me permite conocer el estado de BD */
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == false {
			http.Error(w, "Conexion perdida con la BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
