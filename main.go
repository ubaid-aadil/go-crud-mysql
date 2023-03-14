package main

import (
	"golang-crud-mysql/controllers"
	"golang-crud-mysql/initializers"

	"github.com/gin-gonic/gin"
)

func init() {

	initializers.LoadEnvVariables()
	initializers.ConnectToDatabse()

}

func main() {

	router := gin.Default()

	router.POST("/posts", controllers.PostsCreate)
	router.GET("/posts", controllers.PostsIndex)
	router.GET("/posts/:id", controllers.PostsShow)
	router.PUT("/posts/:id", controllers.PostsUpdate)
	router.DELETE("/posts/:id", controllers.PostDelete)

	router.Run()

}
