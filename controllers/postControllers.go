package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vnaxel/go-jwt/initializers"
	"github.com/vnaxel/go-jwt/models"
)

func CreatePost (c *gin.Context) {
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	userLoggedIn, ok := c.Get("user")

	if !ok {
        log.Printf("Unable to extract user from request context for unknown reason: %v\n", c)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no user in gin request context",
		})

		return
    }

	post := models.Post{Title: body.Title, Body: body.Body, UserID: userLoggedIn.(models.User).ID , User: userLoggedIn.(models.User)}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "an error occured during post publication",
		})

		return
	}

	c.JSON(200, gin.H{
		"post": struct {
			Title string
			Body string
			AuthorID uint
			AuthorEmail string
		}{
			Title: post.Title,
			Body: post.Body,
			AuthorID: post.UserID,
			AuthorEmail: post.User.Email,
		},
	})
}