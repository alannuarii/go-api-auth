package controllers

import (
	"go-api-auth/db"
	"go-api-auth/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
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
		"exp": time.Now().Add(time.Hour).Unix(),
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