package bd

import (
	"context"
	"time"

	"github.com/tikimcrzx/tikimcrzx/models"
)

/*InsertoRelacion graba relación en BD*/
func InsertoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("tikimcrzx")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
