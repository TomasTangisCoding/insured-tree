package app

import (
	"fmt"
	"log"
)

var numLeftChild, numRightChild uint64

type UserTree struct {
	ID           uint64    `json:"id"`
	Name         string    `json:"name"`
	ReferrerID   uint64    `json:"referrer_id"`
	ParentID     uint64    `json:"parent_id"`
	LeftChildID  *UserTree `json:"left_child"`
	RightChildID *UserTree `json:"right_child"`
	IsDelete     bool      `json:"is_delete"`
	CreatedAt    string    `json:"created_at"`
}

type CaculateTree struct {
	NodeID, NumLeftChild, NumRightChild, LeftChildID, RightChildID uint64
}

var userCountMap = make(map[uint64]*CaculateTree)

func FindTreePosition(users []User, referrer_id uint64) (parentID uint64, side string, err error) {

	if len(users) == 0 {
		err := fmt.Errorf("users is empty")
		log.Printf("%v", err)
		return 0, "nil", err
	}
	rootId := users[0].ID
	if rootId == 0 {
		err := fmt.Errorf("invalid rootId")
		log.Printf("%v", err)
		return 0, "nil", err
	}

	userMap := createUserMap(users)

	numLeftChild, numRightChild = 1, 1

	for i := range users {
		rootId = users[i].ID
		root, ok := userMap[rootId]
		if !ok {
			err := fmt.Errorf("user with ID %d not found", rootId)
			log.Printf("%v", err)
			return 0, "nil", err
		}
		numLeftChild, numRightChild = 1, 1
		buildSubtree(root)

		userCountMap[users[i].ID] = &CaculateTree{
			NodeID:        users[i].ID,
			LeftChildID:   users[i].LeftChildID,
			RightChildID:  users[i].RightChildID,
			NumLeftChild:  numLeftChild,
			NumRightChild: numRightChild,
		}
	}

	nodeID, childSide := countSubTree(users[0].ID)
	return nodeID, childSide, nil
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

	return []*UserTree{root}, nil
}

func createUserMap(users []User) map[uint64]*UserTree {
	userMap := make(map[uint64]*UserTree, len(users))
	for i := range users {
		userMap[users[i].ID] = &UserTree{
			ID:         users[i].ID,
			Name:       users[i].Name,
			ReferrerID: users[i].ReferrerID,
			ParentID:   users[i].ParentID,
			IsDelete:   users[i].IsDelete,
			CreatedAt:  users[i].CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	for _, user := range users {
		if user.LeftChildID != 0 {
			userMap[user.ID].LeftChildID = userMap[user.LeftChildID]
		}
		if user.RightChildID != 0 {
			userMap[user.ID].RightChildID = userMap[user.RightChildID]
		}
	}
	return userMap
}

func buildSubtree(node *UserTree) {
	if node.LeftChildID != nil {
		buildSubtree(node.LeftChildID)
		numLeftChild++
	}

	if node.RightChildID != nil {
		buildSubtree(node.RightChildID)
		numRightChild++
	}
}

func countSubTree(userId uint64) (uint64, string) {

	if userCountMap[userId].NumLeftChild < userCountMap[userId].NumRightChild && userCountMap[userId].LeftChildID != 0 {
		return countSubTree(userCountMap[userId].LeftChildID)
	} else if userCountMap[userId].NumLeftChild >= userCountMap[userId].NumRightChild && userCountMap[userId].RightChildID != 0 {
		return countSubTree(userCountMap[userId].RightChildID)
	} else if userCountMap[userId].NumLeftChild == 1 && userCountMap[userId].NumRightChild == 1 {
		return userCountMap[userId].NodeID, "left"
	}
	return userCountMap[userId].NodeID, "right"

}
