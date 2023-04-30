package app

import (
	"log"
)

type UserService struct {
	userModel User
}

func NewUserService(userModel User) *UserService {
	return &UserService{
		userModel: userModel,
	}
}

func (u *UserService) GetUserTree(query string) ([]*UserTree, error) {
	users, err := u.userModel.GetUsers(query)
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

func (u *UserService) SearchUser(query string) ([]User, error) {
	users, err := u.userModel.SearchUser(query)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	return users, nil
}

// // CreateUser creates a new user in the user model
// func (u *UserService) CreateUser(name string) (User, error) {
// 	user := User{Name: name}             // create a new user model with the given name
// 	err := u.userModel.CreateUser(&user) // call the user model method to create a user
// 	if err != nil {
// 		return User{}, err // return an empty user and the error if any
// 	}
// 	return user, nil // return the created user and no error
// }
