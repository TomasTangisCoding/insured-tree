package app

import (
	"fmt"
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

func (u *UserService) CreateUser(name string, referrer_id uint64) ([]User, error) {

	users, err := u.userModel.GetUsersUnlimit(referrer_id)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	// usertree, err := BuildUserTree(users)

	parentID, childSide, err := FindTreePosition(users, referrer_id)
	fmt.Printf("parentID: %d, side: %s\n", parentID, childSide)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	newUser := User{
		Name:       name,
		ReferrerID: referrer_id,
		ParentID:   parentID,
	}
	fmt.Println(newUser)
	createdUser, err := u.userModel.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	oldUser := User{
		ID:           parentID,
		LeftChildID:  0,
		RightChildID: 0,
	}

	if childSide == "left" {
		oldUser.LeftChildID = createdUser[0].ID
	} else if childSide == "right" {
		oldUser.RightChildID = createdUser[0].ID
	}

	fmt.Println(oldUser, createdUser[0].ID)
	updatedUser, err := u.userModel.UpdateUser(oldUser)
	if err != nil {
		return nil, err
	}

	newUsers := []User{}
	newUsers = append(newUsers, createdUser[0], updatedUser[0])

	return newUsers, err
}
