package app

import (
	"fmt"
	"insured/initiate"
	"log"
	"time"
)

type User struct {
	ID           uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string    `gorm:"not null" json:"name"`
	Email        string    `json:"email"`
	ReferrerID   uint64    `json:"referrer_id"`
	ParentID     uint64    `json:"parent_id"`
	LeftChildID  uint64    `json:"left_child_id"`
	RightChildID uint64    `json:"right_child_id"`
	IsDelete     bool      `gorm:"default:false;not null" json:"is_delete"`
	CreatedAt    time.Time `gorm:"type:datetime;not null" json:"created_at"`
}

func (u *User) GetUsers(queryId string) ([]User, error) {
	id := queryId
	const limit = 15
	const query = `
		WITH RECURSIVE cte AS (
			SELECT id, name, left_child_id, right_child_id, parent_id, referrer_id, created_at
			FROM users
			WHERE id = ?
			UNION ALL
			SELECT u.id, u.name, u.left_child_id, u.right_child_id, u.parent_id, u.referrer_id, u.created_at
			FROM users AS u
			JOIN cte ON cte.left_child_id = u.id OR cte.right_child_id = u.id
		)
		SELECT cte.id, cte.name, cte.left_child_id, cte.right_child_id, cte.parent_id, cte.referrer_id, cte.created_at
		FROM cte
		LIMIT ?`

	var users []User
	result := initiate.DBconnect.Raw(query, id, limit).Scan(&users)
	if result.Error != nil {
		log.Printf("%v", result.Error)
		return nil, fmt.Errorf("failed to get users")
	}

	return users, nil
}

func (u *User) SearchUser(query string) ([]User, error) {
	users := []User{}
	result := initiate.DBconnect.Where("ID", query).Find(&users)
	if result.Error != nil {
		log.Printf("%v", result.Error)
		return nil, result.Error
	}
	return users, nil
}
