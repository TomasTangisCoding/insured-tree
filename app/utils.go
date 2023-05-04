package app

import (
	"fmt"
	"log"
)

var numLeftChild, numRightChild uint64
var Ch chan CaculateTree

type UserTree struct {
	ID uint64 `json:"id"`
	// Name       string    `json:"name"`
	// ReferrerID uint64 `json:"referrer_id"`
	// ParentID   uint64    `json:"parent_id"`
	LeftChildID  *UserTree `json:"left_child"`
	RightChildID *UserTree `json:"right_child"`
	// IsDelete   bool      `json:"is_delete"`
	// CreatedAt  string    `json:"created_at"`
}

type CaculateTree struct {
	NodeID, NumLeftChild, NumRightChild, LeftChildID, RightChildID uint64
}

var userCountMap = make(map[uint64]*CaculateTree)

func BuildUserTree(users []User) ([]*UserTree, error) {
	// fmt.Println(users)
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

	//print userMap
	// fmt.Println(userMap)
	// numLeftChild, numRightChild = 1, 1

	// buildSubtree(root)

	// numLeftChild, numRightChild = 1, 1
	// put different rootId into userMap and countSubtree, and get the rootId and numLeftChild and numRightChild
	// then compare them to know where to put a new user
	// userCountMap := make(map[uint64]*CaculateTree, len(users))
	for i := range users {
		rootId = users[i].ID
		root, ok = userMap[rootId]
		if !ok {
			err := fmt.Errorf("user with ID %d not found", rootId)
			log.Printf("%v", err)
			return nil, err
		}
		numLeftChild, numRightChild = 1, 1
		buildSubtree(root)
		if rootId == 3 {
			fmt.Printf("rootId=%v, LeftChild=%v, RightChild=%v\n", root.ID, root.LeftChildID, root.RightChildID)
		}
		fmt.Printf("rootId=%v, numLeftChild=%v, numRightChild=%v\n", rootId, numLeftChild, numRightChild)
		userCountMap[users[i].ID] = &CaculateTree{
			NodeID:        users[i].ID,
			LeftChildID:   users[i].LeftChildID,
			RightChildID:  users[i].RightChildID,
			NumLeftChild:  numLeftChild,
			NumRightChild: numRightChild,
		}
	}
	// fmt.Println(userCountMap)
	for key, value := range userCountMap {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	}

	nodeID, childSide := countSubTree(2)
	fmt.Printf("nodeID=%v, childSide=%v\n", nodeID, childSide)
	// parent_id = nodeID
	// if childSide == "left" {
	// 	LeftChildID = newUser
	// } else {
	// 	RightChildID = newUser
	// }

	return []*UserTree{root}, nil
}

func createUserMap(users []User) map[uint64]*UserTree {
	userMap := make(map[uint64]*UserTree, len(users))
	for i := range users {
		userMap[users[i].ID] = &UserTree{
			ID: users[i].ID,
			// Name:       users[i].Name,
			// ReferrerID: users[i].ReferrerID,
			// ParentID:   users[i].ParentID,
			// IsDelete:   users[i].IsDelete,
			// CreatedAt:  users[i].CreatedAt.Format("2006-01-02 15:04:05"),
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
		fmt.Printf("node=%v, numLeftChild=%v, numRightChild=%v\n", node.ID, numLeftChild, numRightChild)
	}

	if node.RightChildID != nil {
		buildSubtree(node.RightChildID)
		numRightChild++
		fmt.Printf("node=%v, numLeftChild=%v, numRightChild=%v\n", node.ID, numLeftChild, numRightChild)
	}
	fmt.Printf("node=%v, numLeftChild=%v, numRightChild=%v\n", node.ID, numLeftChild, numRightChild)
}

// func countSubTree(userId uint64) (uint64, string) {
// 	fmt.Println("1", userCountMap[userId].NodeID, userCountMap[userId].LeftChildID, userCountMap[userId].RightChildID, userCountMap[userId].NumLeftChild, userCountMap[userId].NumRightChild)

// 	if userCountMap[userId].LeftChildID == 0 || userCountMap[userId].RightChildID == 0 {
// 		if userCountMap[userId].NumLeftChild == 1 {
// 			return userCountMap[userId].NodeID, "left"
// 		} else if userCountMap[userId].NumRightChild == 1 {
// 			return userCountMap[userId].NodeID, "right"
// 		}
// 	}

// 	if userCountMap[userId].RightChildID != 0 && userCountMap[userId].NumLeftChild > userCountMap[userId].NumRightChild {
// 		fmt.Println("3", userCountMap[userId].NodeID, userCountMap[userId].LeftChildID, userCountMap[userId].RightChildID, userCountMap[userId].NumLeftChild, userCountMap[userId].NumRightChild)

// 		countSubTree(userCountMap[userId].RightChildID)

// 		// if userCountMap[userId].NumLeftChild == 1 {
// 		// 	return userCountMap[userId].NodeID, "left"
// 		// } else if userCountMap[userId].NumRightChild == 1 {
// 		// 	return userCountMap[userId].NodeID, "right"
// 		// }

// 	} else if userCountMap[userId].LeftChildID != 0 && userCountMap[userId].NumLeftChild < userCountMap[userId].NumRightChild {
// 		fmt.Println("4", userCountMap[userId].NodeID, userCountMap[userId].LeftChildID, userCountMap[userId].RightChildID, userCountMap[userId].NumLeftChild, userCountMap[userId].NumRightChild)

// 		countSubTree(userCountMap[userId].LeftChildID)

// 		// if userCountMap[userId].NumLeftChild == 1 {
// 		// 	return userCountMap[userId].NodeID, "left"
// 		// } else if userCountMap[userId].NumRightChild == 1 {
// 		// 	return userCountMap[userId].NodeID, "right"
// 		// }
// 	}

// 	fmt.Println("2", userCountMap[userId].NodeID, userCountMap[userId].LeftChildID, userCountMap[userId].RightChildID, userCountMap[userId].NumLeftChild, userCountMap[userId].NumRightChild)

// 	return 999, "right"
// }

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
