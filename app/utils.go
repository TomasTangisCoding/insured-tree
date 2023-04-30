package app

import (
	"fmt"
	"log"
)

type UserTree struct {
	ID         uint64    `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	ReferrerID uint64    `json:"referrer_id"`
	ParentID   uint64    `json:"parent_id"`
	LeftChild  *UserTree `json:"left_child"`
	RightChild *UserTree `json:"right_child"`
	IsDelete   bool      `json:"is_delete"`
	CreatedAt  string    `json:"created_at"`
}

func BuildUserTree(users []User) ([]*UserTree, error) {
	if len(users) == 0 {
		err := fmt.Errorf("users is empty")
		log.Printf("%v", err)
		return nil, err
	}
	rootId := users[0].ID
	if rootId == 0 {
		err := fmt.Errorf("invalid rootId")
		log.Printf("%v", err)
		return nil, err
	}
	root := getUserById(rootId, users)
	if root == nil {
		err := fmt.Errorf("user with ID %d not found", rootId)
		log.Printf("%v", err)
		return nil, err
	}

	buildSubtree(root, users)

	return pointerToSlice(root), nil
}

func buildSubtree(node *UserTree, users []User) {

	if node.LeftChild != nil {
		leftChild := getUserById(node.LeftChild.ID, users)
		if leftChild != nil {
			node.LeftChild = leftChild
			buildSubtree(leftChild, users)
		}
	}

	if node.RightChild != nil {
		rightChild := getUserById(node.RightChild.ID, users)
		if rightChild != nil {
			node.RightChild = rightChild
			buildSubtree(rightChild, users)
		}
	}
}

func getUserById(id uint64, users []User) *UserTree {
	for _, user := range users {
		if user.ID == id {
			return &UserTree{
				ID:         user.ID,
				Name:       user.Name,
				Email:      user.Email,
				ReferrerID: user.ReferrerID,
				ParentID:   user.ParentID,
				IsDelete:   user.IsDelete,
				LeftChild:  getUserById(user.LeftChild, users),
				RightChild: getUserById(user.RightChild, users),
				CreatedAt:  user.CreatedAt.Format("2006-01-02 15:04:05"),
			}
		}
	}
	return nil
}

func pointerToSlice(pointer *UserTree) []*UserTree {
	var slice []*UserTree
	slice = append(slice, pointer)
	return slice
}
