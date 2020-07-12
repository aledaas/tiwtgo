package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/aledaas/tiwtgo/bd"
	"github.com/aledaas/tiwtgo/models"
)

/*GraboTweet graba el tweet en la BD*/
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrio un error al instertar el tweet "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Ocurrio un error al instertar el tweet ", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
