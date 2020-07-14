package bd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aledaas/tiwtgo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*LeoUsuariosTodos Lista los usuarios*/
func LeoUsuariosTodos(ID string, pagina int64, search string, tipo string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("tiwtgo")
	col := db.Collection("usuarios")

	var results []*models.Usuario

	findOptions := options.Find()
	findOptions.SetSkip((pagina - 1) * 20)
	findOptions.SetLimit(20)

	condicion := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := col.Find(ctx, condicion, findOptions)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	var encontrado, incluir bool

	for cursor.Next(ctx) {
		var registro models.Usuario
		err := cursor.Decode(&registro)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = registro.ID.Hex()

		incluir = false

		encontrado, err = ConsultoRelacion(r)
		if tipo == "new" && encontrado == false {
			incluir = true
		}
		if tipo == "follow" && encontrado == true {
			incluir = true
		}
		if r.UsuarioRelacionID == ID {
			incluir = false
		}
		if incluir == true {
			registro.Password = ""
			registro.Biografia = ""
			registro.SitioWeb = ""
			registro.Ubicacion = ""
			registro.Banner = ""
			registro.Email = ""

			results = append(results, &registro)

		}
	}

	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	cursor.Close(ctx)
	return results, true
}
