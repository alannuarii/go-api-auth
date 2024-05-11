package controllers

import (
	"go-api-auth/db"
	"go-api-auth/models"
	"github.com/gin-gonic/gin"
	
)

func GetAllUser(c *gin.Context){
	db := db.DB

	users := []models.User{}

	query := `SELECT * FROM user`

	err := db.Select(&users, query)
	if err != nil{
		c.JSON(500, gin.H{"message": err.Error()})
        return
	}

	c.JSON(200, gin.H{"message": "Sukses", "data": users})
}