package bd

import (
	"context"
	"time"

	"github.com/aledaas/tiwtgo/src/github.com/aledaas/tiwtgo/models"
)

/*BorroRelacion Ejecuta el borrado de una relacion */
func BorroRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("tiwtgo")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
