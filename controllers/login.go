package controllers

import (
	"go-api-auth/db"
	"go-api-auth/models"
	"log"
	"os"
	"time"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	db := db.DB

	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")

	user := models.User{}

	query := `SELECT * FROM user WHERE email = ?`
	err := db.Get(&user, query, email)
	if err != nil {
		c.JSON(404, gin.H{"message": "Email tidak ditemukan"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(401, gin.H{"message": "Password tidak sesuai"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id_user": user.ID,
		"name": user.Name,
		"email": user.Email,
		"role": user.Role,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		c.JSON(500, gin.H{"message": "Internal server error"})
		return
	}

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"message": "Internal server error"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Login berhasil",
		"access_token": tokenString,
	})
}


func VerifyToken(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		c.JSON(401, gin.H{"message": "Access token tidak diberikan"})
		return
	}

	// Hapus prefiks 'Bearer ' jika ada
	accessToken = strings.TrimPrefix(accessToken, "Bearer ")

	// Ambil secret key dari environment variable
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		c.JSON(500, gin.H{"message": "Secret key tidak ditemukan"})
		return
	}

	// Parse access token
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		// Verifikasi metode penandatanganan
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		log.Println(err)
		c.JSON(401, gin.H{"message": "Token tidak valid"})
		return
	}

	// Periksa apakah token valid
	if !token.Valid {
		c.JSON(401, gin.H{"message": "Token tidak valid"})
		return
	}

	// Periksa klaim token jika perlu
	_, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(500, gin.H{"message": "Internal server error"})
		return
	}

	// Anda dapat melakukan sesuatu dengan klaim token di sini jika diperlukan
	// Misalnya, dapat mengambil ID pengguna dari klaim untuk digunakan di aplikasi

	// Jika token valid, beri respons 200 OK
	c.JSON(200, gin.H{"message": "Token valid"})
}