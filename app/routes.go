package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register() {
	r := gin.Default()

	api := r.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := api.Group("/user")

	userController := new(UserController) // Create a new instance of the UserController

	v1.GET("/tree/:id", userController.GetUserTree) // To show original tree, Call the GetUserTree method of the userController
	v1.GET("/:id", userController.SearchUser)       // Call the SearchUser method of the userController
	// v1.POST("/users", userController.CreateUser)        // Call the CreateUser method of the userController

	r.Run(":8080")
}
