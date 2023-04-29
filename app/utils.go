package app

import (
	"fmt"
	"log"
)

func BuildUserTree(users []User) ([]*UserTree, error) {
	if len(users) == 0 { // check if users is empty
		err := fmt.Errorf("users is empty")
		log.Printf("%v", err)
		return nil, err
	}
	rootId := users[0].ID // assume the first user is the root
	if rootId == 0 {      // check if rootId is valid
		err := fmt.Errorf("invalid rootId")
		log.Printf("%v", err)
		return nil, err
	}
	root := getUserById(rootId, users) // get the root pointer
	if root == nil {                   // check if root is nil
		err := fmt.Errorf("user with ID %d not found", rootId)
		log.Printf("%v", err)
		return nil, err
	}

	buildSubtree(root, users)
	//return root as []UserTree

	return pointerToSlice(root), nil
}

func buildSubtree(node *UserTree, users []User) {

	// Traverse left subtree
	if node.LeftChild != nil {
		leftChild := getUserById(node.LeftChild.ID, users)
		if leftChild != nil {
			node.LeftChild = leftChild
			buildSubtree(leftChild, users)
		}
	}

	// Traverse right subtree
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
				LeftChild:  getUserById(user.LeftChild, users),  // use getUserById to get the left child pointer
				RightChild: getUserById(user.RightChild, users), // use getUserById to get the right child pointer
			}
		}
	}
	return nil
}

func pointerToSlice(pointer *UserTree) []*UserTree {
	var slice []*UserTree          // create an empty slice of pointers
	slice = append(slice, pointer) // append the pointer to the slice
	return slice                   // return the final slice
}
