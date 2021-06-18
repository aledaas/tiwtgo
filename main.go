package main

import (
	"log"

	"github.com/aledaas/tiwtgo/src/github.com/aledaas/tiwtgo/bd"
	"github.com/aledaas/tiwtgo/src/github.com/aledaas/tiwtgo/handlers"
)

func main() {

	if bd.ChequeoConnection() == false {
		log.Fatal("sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
