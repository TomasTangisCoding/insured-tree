package app

import (
	"github.com/gin-gonic/gin"
)

func Register() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	api := r.Group("/api")
	routeUser := api.Group("/user")

	userController := new(UserController)

	routeUser.GET("/tree/:id", userController.GetUserTree)
	routeUser.GET("/:id", userController.SearchUser)

	r.Run(":8080")
}
