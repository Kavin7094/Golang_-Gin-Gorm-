package controllers

import (
	models "Learn_GO/Models"
	"Learn_GO/initilaizers"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	// Define request body struct with exported fields + JSON tags
	var requestBody struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	// Bind JSON into requestBody
	// if err := c.ShouldBindJSON(&requestBody); err != nil {
	// 	c.JSON(400, gin.H{"error": err.Error()})
	// 	return
	// }
	c.Bind(&requestBody)
	// Create Post model from requestBody
	post := models.Post{
		Title: requestBody.Title,
		Body:  requestBody.Body}
	// Ensure DB is initialized
	if initilaizers.Db == nil {
		c.JSON(500, gin.H{"error": "Database not initialized"})
		return
	}
	// Insert into DB
	result := initilaizers.Db.Create(&post)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	// Success response
	c.JSON(200, gin.H{"post": post})
}

func PostsIndex(c *gin.Context) {
	var post []models.Post
	initilaizers.Db.Find(&post)
	c.JSON(200, gin.H{"post": post})
}

func PostScope(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	initilaizers.Db.First(&post, id)

	c.JSON(200, gin.H{"post": post})
}

func PostUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	var post models.Post
	initilaizers.Db.First(&post, id)

	initilaizers.Db.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	id := c.Param("id")

	initilaizers.Db.Delete(&models.Post{}, id)

	c.Status(200)
}
