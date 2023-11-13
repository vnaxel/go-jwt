package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vnaxel/go-jwt/controllers"
	"github.com/vnaxel/go-jwt/initializers"
	"github.com/vnaxel/go-jwt/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.POST("/posts", middleware.RequireAuth, controllers.CreatePost)
	r.GET("/posts", middleware.RequireAuth, controllers.GetPosts)

	r.Run() // listen and serve on 0.0.0.0:3000
}