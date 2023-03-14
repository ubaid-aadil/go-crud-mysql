package controllers

import (
	"golang-crud-mysql/initializers"
	"golang-crud-mysql/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	//get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//create a post

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {

		c.Status(400)
		// fmt.Println("error while creating db")
		// fmt.Println(result.Error)
		return

	}
	// return it

	c.JSON(200, gin.H{

		"post": post,
	})

}

func PostsIndex(c *gin.Context) {

	// get the posts

	var posts []models.Post

	initializers.DB.Find(&posts)

	// respond with them

	c.JSON(200, gin.H{

		"posts": posts,
	})

}

func PostsShow(c *gin.Context) {

	//get id of url

	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	// respond with data

	c.JSON(200, gin.H{

		"post": post,
	})

}

func PostsUpdate(c *gin.Context) {

	// get the id of the url

	id := c.Param("id")

	//get the data of the req body

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	//find the post we have  to update
	var post models.Post
	initializers.DB.First(&post, id)

	//update it

	initializers.DB.Model(&post).Updates(models.Post{

		Title: body.Title,
		Body:  body.Body,
	})

	//respond with the data

	c.JSON(200, gin.H{

		"post": post,
	})

}

func PostDelete(c *gin.Context) {

	// Get the id of the url

	id := c.Param("id")

	//delete the post

	initializers.DB.Delete(&models.Post{}, id)

	//respond

	c.Status(200)

}
