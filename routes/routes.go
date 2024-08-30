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
	router.GET("/user/:username", handlers.GetPublicProfile) // get public profile of a user

	// protected routes

	// user routes
	router.GET("/profile", handlers.AuthMiddleware(), handlers.GetProfile)               // get own profile
	router.PATCH("/profile", handlers.AuthMiddleware(), handlers.UpdateProfile)          // update basic profile info
	router.DELETE("/profile/details", handlers.AuthMiddleware(), handlers.DeleteProfile) // delete
	router.GET("/profile/details", handlers.AuthMiddleware(), handlers.GetProfileDetails)
	//router.PATCH("profiles/details", handlers.AuthMiddleware(), handlers.UpdateProfileDetails) // get email and other private info
	// @TODO /profile/details (PUT) - update email and other private info

	// admin routes

	// Start the server
	router.Run(":8080")
}
