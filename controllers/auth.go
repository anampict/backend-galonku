package controllers

import (
	"backendGalonku/config"
	"backendGalonku/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterUser(c *gin.Context) {
	var input models.User

	// ambil json dari request dan bind ke struct User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "data tidak valid"})
		return
	}

	//cek apakah email sudah terdaftar
	userCollection := config.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var existingUser models.User
	err := userCollection.FindOne(ctx, gin.H{"email": input.Email}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email sudah terdaftar"})
		return
	}

	// buat user baru
	newUser := models.User{
		ID:       primitive.NewObjectID(),
		Email:    input.Email,
		Password: input.Password, // pastikan password di-hash sebelum disimpan
		Role:     "user", // default role
	}

	_, err = userCollection.InsertOne(ctx, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal membuat user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user berhasil dibuat", "user": newUser})

}


