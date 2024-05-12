package controllers

import (
	"go-api-auth/db"
	"go-api-auth/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InsertUser(c *gin.Context){
	db := db.DB

	name := c.Request.FormValue("name")
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")
	role := c.Request.FormValue("role")

	hashPassword := utils.SetPassword(password)

	query := `INSERT INTO user (name, email, password, role) VALUES (?, ?, ?, ?)`

	_, err := db.Exec(query, name, email, hashPassword, role)
	if err != nil{
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": fmt.Sprintf("Berhasil menambahkan user %s", name)})
}