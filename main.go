package main

import (
	"github.com/gin-gonic/gin"
	"go-api-auth/controllers"
)

func main(){
	r:= gin.Default()

	r.GET("api/users", controllers.GetAllUser)
	r.GET("api/user/:id", controllers.GetUser)

	r.POST("api/user", controllers.InsertUser)

	r.PUT("api/user/:id", controllers.UpdateUser)

	r.DELETE("api/user/:id", controllers.DeleteUSer)

	r.Run(":8888")
}