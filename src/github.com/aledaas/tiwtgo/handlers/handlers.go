package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/aledaas/tiwtgo/middlew"
	"github.com/aledaas/tiwtgo/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores pongo a escuchar al servidor */
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
