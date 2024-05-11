package main

import (
	"github.com/gin-gonic/gin"
	"go-api-auth/controllers"
)

func main(){
	r:= gin.Default()

	r.GET("api/alluser", controllers.GetAllUser)

	r.Run(":8888")
}