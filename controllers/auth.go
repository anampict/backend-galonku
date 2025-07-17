package controllers

import (
	"backendGalonku/config"
	"backendGalonku/models"
	"backendGalonku/utils"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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

	// hash password sebelum disimpan
	hashedPassword, err := utils.HashPassword(input.Password)
	if  err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mengenkripsi password"})
		return
		
	}

	// buat user baru
	newUser := models.User{
		ID:       primitive.NewObjectID(),
		Email:    input.Email,
		Password: hashedPassword, // pastikan password di-hash sebelum disimpan
		Role:     "user", // default role
	}

	_, err = userCollection.InsertOne(ctx, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal membuat user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user berhasil dibuat", "user": newUser})

}


//login

func LoginUser(c *gin.Context) {
	var input models.User // Buat variabel struct untuk menyimpan input dari user

	// Ambil JSON dari request dan masukkan ke struct 'input'
	if  err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "data tidak valid"})
		return
		
	}

	// Akses collection 'users' dari database
	userCollection := config.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Buat context timeout 5 detik agar koneksi database tidak menggantung
	defer cancel()

	var  user models.User // Variabel untuk menyimpan data user dari database
	err := userCollection.FindOne(ctx,bson.M{"email": input.Email}).Decode(&user)// Cari user berdasarkan email yang diinput 
	 // Kalau user tidak ditemukan, kirim pesan error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "email tidak ditemukan"})
		return
	}

	// Bandingkan password input dengan hash di database
	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "password salah"})
		return
	}

	// Buat JWT token jika login sukses
	token, err := utils.GenerateToken(user.ID.Hex(), user.Role)

	// Jika ada error saat membuat token, kirim pesan error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal membuat token"})
		return
	}

	 // Kirim response ke user berupa token dan role
	c.JSON(http.StatusOK, gin.H{"message": "login berhasil", "token": token, "user": user})
}


