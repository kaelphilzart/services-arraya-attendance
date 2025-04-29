package interfaces

import (
	"encoding/json"
	"time"

	uuid "github.com/google/uuid" 
)

type UserProfile struct {
	ID			uuid.UUID 				`db:"id, primarykey" json:"id"`
	User		*json.RawMessage		`db:"user" json:"user"`
	FullName	string					`db:"full_name" json:"full_name"`
	BirthDate   string      			`db:"birth_date" json:"birth_date"`
	BirthPlace	string					`db:"birth_place" json:"birth_place"`
	Address     string      			`db:"address" json:"address"`
	PhoneNumber	string					`db:"phone_number" json:"phone_number"`
	Gender      string      			`db:"gender" json:"gender"`
	Photo		string					`db:"photo" json:"photo"`
	UpdatedAt   *time.Time       		`db:"updated_at" json:"updated_at"`
	CreatedAt   time.Time        		`db:"created_at" json:"created_at"`
}