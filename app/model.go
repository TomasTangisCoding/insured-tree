package app

import (
	"fmt"
	"insured/db"
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
}

type User struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string `gorm:"not null" json:"name"`
	Email      string `json:"email"`
	ReferrerID uint64 `json:"referrer_id"`
	ParentID   uint64 `json:"parent_id"`
	LeftChild  uint64 `json:"left_child"`
	RightChild uint64 `json:"right_child"`
	IsDelete   bool   `gorm:"default:false;not null" json:"is_delete"`
	// CreatedAt  time.Time  `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	// UpdatedAt  *time.Time `gorm:"type:datetime" json:"updated_at"`
}

// func (u User) MarshalJSON() ([]byte, error) {
// 	type Alias User
// 	return json.Marshal(&struct {
// 		CreatedAt string `json:"created_at"`
// 		*Alias
// 	}{
// 		CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
// 		Alias:     (*Alias)(&u),
// 	})
// }

func (u *User) GetUserTree(queryId string) ([]User, error) {
	id := queryId
	const limit = 15
	const query = `
		WITH RECURSIVE cte AS (
			SELECT id, name, left_child, right_child, parent_id, referrer_id
			FROM users
			WHERE id = ?
			UNION ALL
			SELECT u.id, u.name, u.left_child, u.right_child, u.parent_id, u.referrer_id
			FROM users AS u
			JOIN cte ON cte.left_child = u.id OR cte.right_child = u.id
		)
		SELECT cte.id, cte.name, cte.left_child, cte.right_child, cte.parent_id, cte.referrer_id
		FROM cte
		LIMIT ?`

	var users []User
	result := db.DBconnect.Raw(query, id, limit).Scan(&users)
	if result.Error != nil {
		log.Printf("%v", result.Error)
		return nil, fmt.Errorf("failed to get users")
	}

	return users, nil
}

func (u *User) SearchUser(query string) ([]User, error) {
	users := []User{}
	result := db.DBconnect.Where("ID", query).Find(&users)
	if result.Error != nil {
		log.Printf("%v", result.Error)
		return nil, result.Error
	}
	return users, nil
}
