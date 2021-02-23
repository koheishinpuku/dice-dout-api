package models

import (
	"time"

	d "github.com/koheishinpuku/dice-dout-api/db"
)

// User :
type User struct {
	ID         string    `json:"user_id"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
	DeletedAt  time.Time `json:"-"`
	Name       string    `json:"name"`
	HavePicies int       `json:"have_picies"`
	PicePlace  int       `json:"pice_place"`
	GoalPicies int       `json:"goal_picies"`
}

// GetUser :
func GetUser(userID string) *User {
	var user User
	d.DB.Where("id = ?", userID).Find(&user)
	return &user
}
