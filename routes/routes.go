package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"profiles_go/handlers"
)

// this file is to be used in main.go

func Routes() {
	// Set up the router
	router := gin.Default()

	router.Use(cors.Default()) // enable cors

	// public routes

	public := router.Group("/api/v1")
	{
		public.POST("/register", handlers.Register)
		public.POST("/login", handlers.Login)
		public.GET("/user/:username", handlers.GetPublicProfile) // get public profile of a user
	}

	// protected routes
	protected := router.Group("/api/v1/profile")
	protected.Use(handlers.AuthMiddleware())
	{
		// user routes
		protected.GET("/", handlers.AuthMiddleware(), handlers.GetProfile)              // get own profile
		protected.PATCH("/", handlers.AuthMiddleware(), handlers.UpdateProfile)         // update basic profile info
		protected.DELETE("/details", handlers.AuthMiddleware(), handlers.DeleteProfile) // delete
		protected.GET("/details", handlers.AuthMiddleware(), handlers.GetProfileDetails)
		//router.PATCH("profiles/details", handlers.AuthMiddleware(), handlers.UpdateProfileDetails) // get email and other private info
		// @TODO /profile/details (PUT) - update email and other private info
	}

	// admin routes

	// Start the server
	router.Run(":8080")
}
