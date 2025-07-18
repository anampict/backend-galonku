package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type Transaction struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      primitive.ObjectID `bson:"user_id" json:"user_id"` // ID pengguna yang melakukan transaksi
	ProductID   primitive.ObjectID `bson:"product_id" json:"product_id"` // ID produk yang dibeli
	Quantity    int                `bson:"quantity" json:"quantity"` // Jumlah produk yang dibeli
	TotalPrice  float64            `bson:"total_price" json:"total_price"` // Total harga transaksi
	TransactionDate primitive.DateTime    `bson:"transaction_date" json:"transaction_date"` // Tanggal transaksi
	Status      string             `bson:"status" json:"status"` // Status transaksi (misalnya: "pending", "completed", "cancelled")
}