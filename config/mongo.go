// koneksi ke database MongoDB
package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	mongoURI := os.Getenv("MONGODB_URI") //  Ambil dari environment variable
	if mongoURI == "" {
		log.Fatal("Environment variable MONGODB_URI belum diatur")
	}

	clientOptions := options.Client().
		ApplyURI(mongoURI)

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