package controllers

import (
	"backendGalonku/config"
	"backendGalonku/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//inisiasi collection dari mongo db
var productCollection *mongo.Collection = config.DB.Collection("products")

//fungsi untuk membuat produk baru

func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	product.ID = primitive.NewObjectID() //generate id baru
	product.CreatedAt = time.Now() //set waktu dibuat

	//simpan ke database
	_, err := productCollection.InsertOne(c.Request.Context(), product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal menambahkan produk"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Produk berhasil ditambahkan", "product": product})

}