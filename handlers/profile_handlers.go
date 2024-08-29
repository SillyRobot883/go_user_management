package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"profiles_go/db"
)

// this file uses profile_repo.go to interact with the database

// get own profile
func GetProfile(c *gin.Context) {
	username := c.MustGet("username").(string)
	user, err := db.GetProfileByUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// this retrieves the private profile details of a user
func GetProfileDetails(c *gin.Context) {
	username := c.MustGet("username").(string)
	user, err := db.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// this retrieves the public profile details of a user
// used by other users
func GetPublicProfile(c *gin.Context) {
	username := c.Param("username")
	profile, err := db.GetPublicProfileByUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error getting public profile": err.Error()})
		return
	}

	if profile == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	c.JSON(http.StatusOK, profile)
}

// UpdateProfile updates the profile of the user based on the context username
func UpdateProfile(c *gin.Context) {
	username := c.MustGet("username").(string)
	var input db.UpdateProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := db.UpdateProfileByUsername(username, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteProfile(c *gin.Context) {
	username := c.MustGet("username").(string)
	err := db.DeleteProfileByUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile deleted successfully"})
}
