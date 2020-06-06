package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/tikimcrzx/tikimcrzx/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*LeoUsuarios Lee todos los usuarios en el sstema si trae R es quien se relaciona con el usuario*/
func LeoUsuarios(ID string, pagina int64, buscar string, tipo string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("tikimcrzx")
	col := db.Collection("usuarios")

	var resultados []*models.Usuario

	findOptions := options.Find()
	findOptions.SetSkip((pagina - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + buscar},
	}

	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return resultados, false
	}

	var encontrado, incluir bool

	for cur.Next(ctx) {
		var s models.Usuario
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return resultados, false
		}

		var r models.Relacion
		r.UsuarioID = ID
		r.UsurioRelacionID = s.ID.Hex()

		incluir = false

		encontrado, err = ConsultoRelacion(r)
		if tipo == "new" && encontrado == false {
			incluir = true
		}
		if tipo == "follow" && encontrado == true {
			incluir = true
		}
		if r.UsurioRelacionID == ID {
			incluir = false
		}
		if incluir == true {
			s.Password = ""
			s.Biografia = ""
			s.SitioWeb = ""
			s.Ubicacion = ""
			s.Banner = ""
			s.Email = ""

			resultados = append(resultados, &s)
		}
	}

	err = cur.Err()
	if err != nil {
		fmt.Print(err.Error())
		return resultados, false
	}
	cur.Close(ctx)
	return resultados, true
}
