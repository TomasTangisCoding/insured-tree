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
