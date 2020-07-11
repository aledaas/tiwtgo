package main

import (
	"log"

	"github.com/aledaas/tiwtgo/bd"
	"github.com/aledaas/tiwtgo/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("son conexion a la BD")
		return
	}
	handlers.Manejadores()
}
