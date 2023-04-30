package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService UserService
}

func NewUserController(userService UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (u *UserController) GetUserTree(c *gin.Context) {
	query := c.Param("id")
	users, err := u.userService.GetUserTree(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("%v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (u *UserController) SearchUser(c *gin.Context) {
	query := c.Param("id")
	users, err := u.userService.SearchUser(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("%v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

// func (u *UserController) CreateUser(c *gin.Context) {
// 	var request struct {
// 		Name string `json:"name" binding:"required"`
// 	}
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	user, err := u.userService.CreateUser(request.Name) // call the user service method to create a user with the given name
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		log.Fatal(err)
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"user": user,
// 	})
// }
