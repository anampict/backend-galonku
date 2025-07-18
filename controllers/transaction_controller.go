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
