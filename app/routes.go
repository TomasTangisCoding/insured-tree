package app

import (
	"github.com/gin-gonic/gin"
)

func Register() {
	r := gin.Default()

	api := r.Group("/api")
	routeUser := api.Group("/user")

	userController := new(UserController)

	routeUser.GET("/tree/:id", userController.GetUserTree)
	routeUser.GET("/:id", userController.SearchUser)
	// routeUser.POST("/users", userController.CreateUser)

	r.Run(":8080")
}
