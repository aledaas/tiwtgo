package routers

import (
	"net/http"

	"github.com/aledaas/tiwtgo/bd"
	"github.com/aledaas/tiwtgo/models"
)

/*AltaRelacion Registra la relacion entre usuarios*/
func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}
	var t models.Relacion

	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al instertar la relacion "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Ocurrio un error al instertar la relacion ", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
