package controllers

import (
	"go-api-auth/db"
	"go-api-auth/models"
	"github.com/gin-gonic/gin"
	
)

func GetAllUser(c *gin.Context){
	db := db.DB

	users := []models.User{}

	query := `SELECT id_user, name, email, role FROM user`

	err := db.Select(&users, query)
	if err != nil{
		c.JSON(500, gin.H{"message": err.Error()})
        return
	}

	c.JSON(200, gin.H{"message": "Sukses", "data": users})
}

func GetUser(c *gin.Context){
	db := db.DB

	userId := c.Param("id")

	user := []models.User{}

	query := `SELECT id_user, name, email, role FROM user WHERE id_user = ?`

	err := db.Select(&user, query, userId)
	if err != nil{
		c.JSON(500, gin.H{"message": err.Error()})
        return
	}

	c.JSON(200, gin.H{"message": "Sukses", "data": user})
}