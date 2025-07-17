package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtkey = []byte ("secret_galonku")

//struktur payload untuk menyimpan data yang akan disimpan di token

type Claims struct {
	UserId string `json:"userId"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// Middleware untuk memeriksa token JWT
func AuthMiddleware() gin.HandlerFunc {
	return  func(ctx *gin.Context) {
		//ambil token dari header Authorization
		authheader := ctx.GetHeader("Authorization")
		if authheader == "" {
			ctx.JSON(401, gin.H{"error": "token tidak ditemukan"})
			ctx.Abort()
			return
		}

		//verifikasi token
		parts := strings.Split(authheader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			 ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Format token salah"})
             ctx.Abort()
             return
		}

		tokenString := parts[1]

		// parsing dan verifikasi token
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtkey, nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			ctx.Abort()
			return
		}

		// Cek apakah token sudah expired
		if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token sudah expired"})
			ctx.Abort()
			return
		}

		// Simpan data user ke context supaya bisa dipakai di handler/controller
		ctx.Set("user_id", claims.UserId)
		ctx.Set("role", claims.Role)

		ctx.Next()
	}
}