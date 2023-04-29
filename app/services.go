package app

import (
	"log"
)

// UserService is a struct that holds the methods for handling user requests
type UserService struct {
	userModel User // inject the user model as a dependency
}

// NewUserService creates and returns a new UserService instance
func NewUserService(userModel User) *UserService {
	return &UserService{
		userModel: userModel,
	}
}

// GetUserTree returns a slice of users from the user model
func (u *UserService) GetUserTree(query string) ([]*UserTree, error) {
	users, err := u.userModel.GetUserTree(query)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	usertree, err := BuildUserTree(users)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	return usertree, nil
}

// SearchUser returns a slice of users that match the given query from the user model
func (u *UserService) SearchUser(query string) ([]User, error) {
	users, err := u.userModel.SearchUser(query) // call the user model method to search users by query
	if err != nil {
		log.Printf("%v", err)
		return nil, err // return the error if any
	}
	return users, nil // return the users if no error
}

// // CreateUser creates a new user in the user model
// func (u *UserService) CreateUser(name string) (User, error) {
// 	user := User{Name: name}                 // create a new user model with the given name
// 	err := u.userModel.CreateUser(&user) // call the user model method to create a user
// 	if err != nil {
// 		return User{}, err // return an empty user and the error if any
// 	}
// 	return user, nil // return the created user and no error
// }
