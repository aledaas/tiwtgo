package routers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/aledaas/tiwtgo/src/github.com/aledaas/tiwtgo/bd"
	"github.com/aledaas/tiwtgo/src/github.com/aledaas/tiwtgo/models"
)

/*SubirBanner se encarga de subir el avatar*/
func SubirBanner(w http.ResponseWriter, r *http.Request) {
	// De la peticion traemos el archivo y el nombre del archivo biene en handler
	file, hanlder, err := r.FormFile("banner")
	if err != nil {
		http.Error(w, "Se ha presentado un error: "+err.Error(), http.StatusBadRequest)
		return
	}
	// Verificar la extension
	var extension = strings.Split(hanlder.Filename, ".")[1]
	if strings.ToLower(extension) == "png" || strings.ToLower(extension) == "jpg" ||
		strings.ToLower(extension) == "gif" {
		// Comprobar si ya existe otra imagen con ese perfil
		comprueba, ext := comprobarBanner()
		fmt.Println(comprueba, ext)
		// Si la funcion comprobar encontro una imagen antetior asociada a ese
		// perfil la elimina
		if comprueba {
			os.Remove("uploads/banners/" + IDUsuario + "." + ext)
		}
		// Carpeta donde se van a subir los archivos, se debe grabar con el id del usuario de la db
		var archivo string = "uploads/banners/" + IDUsuario + "." + extension

		f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
		// Verificar si pudo abrir el archivo
		if err != nil {
			http.Error(w, "No se puede leer el archivo: "+err.Error(), http.StatusBadRequest)
			return
		}
		// Copiar y renombrar
		_, err = io.Copy(f, file)
		if err != nil {
			http.Error(w, "No se pudo copiar el archivo: "+err.Error(), http.StatusBadRequest)
			return
		}

		var usuario models.Usuario
		var status bool

		usuario.Banner = IDUsuario + "." + extension
		status, err = bd.ModificoRegistro(usuario, IDUsuario)

		if err != nil || status == false {
			http.Error(w, "No se actualizar el banner del perdil: "+err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "applicacion/json")
		w.WriteHeader(http.StatusCreated)
	} else {
		http.Error(w, "La extension no es valida ", http.StatusBadRequest)
		return
	}

}

/*comprobarBanner() verificar si un usuario ya tiene una imagen de perfil para eliminarla*/
func comprobarBanner() (bool, string) {
	// Se declaran el tipo de imagenes de a buscar
	var extensiones = []string{"png", "gif", "jpg"}
	// Se itera el slicer
	for i := 0; i < len(extensiones); i++ {
		// Comprobar si ya existe otra imagen con ese perfil para eliminarla
		var archivo = "uploads/banners/" + IDUsuario + "." + extensiones[i]
		// Trata de verificar que existe
		j, err := os.Open(archivo)
		if err != nil {
			j.Close()
			// Si existe la imagen retorna la extension que encontro
		} else {
			return true, extensiones[i]
		}
	}
	// Si no existe aun imagen para ese usuario retorna falso y vacio
	return false, string("")
}
