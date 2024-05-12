package controllers

import (
	"fmt"
	"go-api-auth/db"
	"go-api-auth/utils"
	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context){
	db := db.DB

	userId := c.Param("id")

	name := c.Request.FormValue("name")
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")
	role := c.Request.FormValue("role")

	hashPassword := utils.SetPassword(password)

	queryUpdate := `UPDATE user SET name = ?, email = ?, password = ?, role = ? WHERE id_user = ?`

	_, err := db.Exec(queryUpdate, name, email, hashPassword, role, userId)
	if err != nil{
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": fmt.Sprintf("Berhasil mengupdate user %s", name)})
}