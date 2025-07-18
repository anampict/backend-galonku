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

//inisiasi collection dari mongo db
func getProductCollection() *mongo.Collection {
    collection := config.DB.Collection("products")
    return collection
}


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
	_, err := getProductCollection().InsertOne(c.Request.Context(), product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal menambahkan produk"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Produk berhasil ditambahkan", "product": product})

}

//fungsi untuk mendapatkan semua produk

func GetAllProducts(c *gin.Context) {
	var products[]models.Product

	cursor, err := getProductCollection().Find(context.TODO(), bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mendapatkan produk"})
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal decode produk"})
			return
		}
		products = append(products, product)
	}

	c.JSON(http.StatusOK, gin.H{"message": "berhasil mendapatkan produk", "products": products})
}


// fungsi untuk menghapus produk berdasarkan ID

func DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	_, err = getProductCollection().DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal menghapus produk"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil dihapus"})
}

// fungsi untuk mengupdate produk berdasarkan ID

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var updatedProduct models.Product
	if err := c.BindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"name":        updatedProduct.Name,
			"price":       updatedProduct.Price,
			"stock":       updatedProduct.Stock,
			"description": updatedProduct.Description,
		},
	}

	_, err = getProductCollection().UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mengupdate produk"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil diupdate"})
}