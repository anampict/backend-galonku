package controllers

import (
	"backendGalonku/config"
	"backendGalonku/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Gunakan nama variabel berbeda agar tidak bentrok
func getTransactionCollection() *mongo.Collection {
    return config.DB.Collection("transactions")
}
func getProductCollectiontransaksi() *mongo.Collection {
    return config.DB.Collection("products")
}

func CreateTransaction(c *gin.Context) {
	var req struct {
		UserID    string `json:"user_id"`
		ProductID string `json:"product_id"`
		Quantity  int    `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	productID, err := primitive.ObjectIDFromHex(req.ProductID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	userID, err := primitive.ObjectIDFromHex(req.UserID)
	if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
    return
	}


	// Ambil data produk
	var product models.Product
	err = getProductCollectiontransaksi().FindOne(context.TODO(), bson.M{"_id": productID}).Decode(&product)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	totalPrice := float64(req.Quantity) * product.Price


	transaction := models.Transaction{
		ID:         primitive.NewObjectID(),
		UserID:    userID, // Ganti dengan ID user yang sesuai
		ProductID:  productID,
		Quantity:   req.Quantity,
		TotalPrice: totalPrice,
		TransactionDate: primitive.NewDateTimeFromTime(time.Now()),
    	Status:          "pending",
	}

	_, err = getTransactionCollection().InsertOne(context.TODO(), transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction created", "data": transaction})
}

// Fungsi untuk mendapatkan semua transaksi

func GetAllTransactions(c *gin.Context) {
	cursor, err := getTransactionCollection().Find(context.TODO(), bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal ambil data transaksi"})
		return
	}

	defer cursor.Close(context.TODO())

	var transactions []models.Transaction

	if err = cursor.All(context.TODO(), &transactions); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal decode transaksi"})
		return
		
	}
	c.JSON(http.StatusOK, gin.H{"message": "berhasil mendapatkan transaksi", "transactions": transactions})


}

// fungsi get transaksi berdasarkan ID

func GetTransactionByID(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := primitive.ObjectIDFromHex(userIDStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	filter := bson.M{"user_id": userID}
	cursor,err := getTransactionCollection().Find(context.TODO(), filter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mendapatkan transaksi"})
		return

	}
	defer cursor.Close(context.TODO())

	var transactions []models.Transaction
	if err = cursor.All(context.TODO(), &transactions); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal decode transaksi"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "berhasil mendapatkan transaksi id tersebut", "transactions": transactions})
}

// fungsi untuk mengupdate status transaksi berdasarkan ID

func UpdateTransactionStatus(c *gin.Context) {
	id := c.Param("id")
	transactionID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var req struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid status"})
		return
	}

	update := bson.M{"$set": bson.M{"status": req.Status}}
	_,err = getTransactionCollection().UpdateOne(context.TODO(), bson.M{"_id": transactionID}, update)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mengupdate status transaksi"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Status transaksi berhasil diupdate"})
}
