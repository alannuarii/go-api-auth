package controllers

import (
	"fmt"
	"go-api-auth/db"

	"github.com/gin-gonic/gin"
)

func DeleteUSer(c *gin.Context){
	db := db.DB

	userId := c.Param("id")

	queryGetName := `SELECT name FROM user WHERE id_user = ?`

	var name string
	err := db.QueryRow(queryGetName, userId).Scan(&name)
	if err != nil{
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	queryDelete := `DELETE FROM user WHERE id_user = ?`

	_, err = db.Exec(queryDelete, userId)
	if err != nil{
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": fmt.Sprintf("Berhasil menghapus user %s", name)})
}