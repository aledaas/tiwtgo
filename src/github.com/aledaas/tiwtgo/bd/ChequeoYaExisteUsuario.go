package bd

import (
	"context"
	"time"

	"github.com/aledaas/tiwtgo/src/github.com/aledaas/tiwtgo/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario chequea si existe el usuario */
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("tiwtgo")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
