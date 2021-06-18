package routers

import (
	"encoding/json"
	"net/http"

	"github.com/aledaas/tiwtgo/src/github.com/aledaas/tiwtgo/bd"
	"github.com/aledaas/tiwtgo/src/github.com/aledaas/tiwtgo/models"
)

/*Registro es la funcion para crear en la BD el registro del usuario */
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "El password debe ser mayor a 6 caracteres", 400)
		return
	}
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "El email ya existe en la BD", 400)
		return
	}
	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al almacenar en la BD"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "no se logro insertar el registro del usuario", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
