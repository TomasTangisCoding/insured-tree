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

	userMap := createUserMap(users)

	root, ok := userMap[rootId]
	if !ok {
		err := fmt.Errorf("user with ID %d not found", rootId)
		log.Printf("%v", err)
		return nil, err
	}

	buildSubtree(root)

	tree := []*UserTree{root}
	return tree, nil
}

func createUserMap(users []User) map[uint64]*UserTree {
	userMap := make(map[uint64]*UserTree, len(users))
	for i := range users {
		userMap[users[i].ID] = &UserTree{
			ID:         users[i].ID,
			Name:       users[i].Name,
			Email:      users[i].Email,
			ReferrerID: users[i].ReferrerID,
			ParentID:   users[i].ParentID,
			IsDelete:   users[i].IsDelete,
			CreatedAt:  users[i].CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	for _, user := range users {
		if user.LeftChild != 0 {
			userMap[user.ID].LeftChild = userMap[user.LeftChild]
		}
		if user.RightChild != 0 {
			userMap[user.ID].RightChild = userMap[user.RightChild]
		}
	}
	return userMap
}

func buildSubtree(node *UserTree) {
	if node.LeftChild != nil {
		buildSubtree(node.LeftChild)
	}

	if node.RightChild != nil {
		buildSubtree(node.RightChild)
	}
}
