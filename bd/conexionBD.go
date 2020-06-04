package bd

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConectarBD()

// var clientOptions = options.Client().ApplyURI("mongodb://127.0.0.1:27017/")

/*ConectarBD función para conectar a la base de datos*/
func ConectarBD() *mongo.Client {
	mongoUri := os.Getenv("MONGO_URI")
	if mongoUri == "" {
		mongoUri = "mongodb://127.0.0.1:27017/"
	}
	clientOptions := options.Client().ApplyURI(mongoUri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión Exitosa con la BD")
	return client
}

/*chequeoConexion verifica conexión a la base de datos */
func ChequeoConexion() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
