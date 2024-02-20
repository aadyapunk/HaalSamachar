package main

import (
	"blog/contracttesting/db"
	"blog/contracttesting/models"
	"blog/contracttesting/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getUserByIDHandler(c *gin.Context) {
	// Extract user ID from request parameters
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Perform database operation to get user by ID
	user, err := db.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	// Return user data as JSON response
	c.JSON(http.StatusOK, user)
}

func getUserBlogPostsHandler(c *gin.Context) {
	// Extract user ID from request parameters
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Perform database operation to get blogs by user ID
	blogs, err := db.GetBlogPostsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user blogs"})
		return
	}

	// Return blogs data as JSON response
	c.JSON(http.StatusOK, blogs)
}

func getUserCommentsHandler(c *gin.Context) {
	// Extract user ID from request parameters
	userIDStr := c.Param("id")
	// Perform database operation to get comments by user ID
	comments, err := db.GetCommentsByUserID(userIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user comments"})
		return
	}

	// Return comments data as JSON response
	c.JSON(http.StatusOK, comments)
}

func SignupHandler(c *gin.Context) {
	// Parse JSON request body
	var req models.User
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}
	hashedPassword := utils.HashPassword(req.Password)

	// Create a new user in the database
	user, err := db.CreateUser(req.Username, req.Email, hashedPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Return user data as JSON response
	c.JSON(http.StatusOK, user)
}