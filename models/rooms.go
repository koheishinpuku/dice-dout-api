package models

import (
	"time"

	d "github.com/koheishinpuku/dice-dout-api/db"
)

// Room :
type Room struct {
	ID            string    `json:"user_id"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
	DeletedAt     time.Time `json:"-"`
	Name          string    `json:"name"`
	PlayerOneID   string    `json:"player_one_id"`
	PlayerTwoID   string    `json:"player_two_id"`
	PlayerThreeID string    `json:"player_three_id"`
	PlayerFourID  string    `json:"player_four_id"`
	User          *User
}

// GetRoom :
func GetRoom(roomID string) *Room {
	var room Room
	d.DB.Where("id = ?", roomID).Find(room)

	return &room
}
