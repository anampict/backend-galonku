package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"` //nama produk
	Price      float64            `bson:"price" json:"price"` //harga produk
	Stock       int                `bson:"stock" json:"stock"` //jumlah stok produk
	Description string             `bson:"description" json:"description"` //deskripsi produk
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"` //waktu produk dibuat
}