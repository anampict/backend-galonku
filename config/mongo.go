// koneksi ke database MongoDB
package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	clientOptions := options.Client().
		ApplyURI("mongodb://localhost:27017")

		ctx,cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("gagal connect ke database:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("gagal ping database:", err)
	}

	DB = client.Database("galonku")
	fmt.Println("Berhasil connect ke database MongoDB")
}

