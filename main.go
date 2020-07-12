package main

import (
	"log"

	"github.com/aledaas/tiwtgo/bd"
	"github.com/aledaas/tiwtgo/handlers"
)

func main() {
	if bd.ChequeoConnection() == false {
		log.Fatal("son conexion a la BD")
		return
	}
	handlers.Manejadores()
}
