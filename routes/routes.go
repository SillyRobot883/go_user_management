package routes

import (
	"github.com/gin-gonic/gin"
	"profiles_go/handlers"
)

// this file is to be used in main.go

func Routes() {
	// Set up the router
	router := gin.Default()

	// public routes
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)
	// @TODO - users can access other users' public profiles
	router.GET("/user/:username", handlers.GetPublicProfile) // get public profile of a user

	// protected routes
	router.GET("/profile", handlers.AuthMiddleware(), handlers.GetProfile)                // get own profile
	router.PUT("/profile", handlers.AuthMiddleware(), handlers.UpdateProfile)             // update
	router.DELETE("/profile", handlers.AuthMiddleware(), handlers.DeleteProfile)          // delete
	router.GET("/profile/details", handlers.AuthMiddleware(), handlers.GetProfileDetails) // get email and other private info

	// get other user's profile

	// Start the server
	router.Run(":8080")
}
