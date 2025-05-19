package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConectarDB() {
	env := os.Getenv("ENV")
	var mongoURI string

	if env == "dev" {
		mongoURI = os.Getenv("MONGO_URI_TEST")
		fmt.Println("🧪 Entorno de desarrollo: conectando a Mongo LOCAL")
	} else {
		mongoURI = os.Getenv("MONGO_URI")
		fmt.Println("🚀 Entorno de producción: conectando a Mongo DOCKER")
	}
	
	dbName := os.Getenv("MONGO_DB")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}

	DB = client.Database(dbName)
	fmt.Println("✅ Conectado a MongoDB")
}

func GetCollection(nombre string) *mongo.Collection {
	return DB.Collection(nombre)
}
