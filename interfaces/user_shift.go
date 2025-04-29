package interfaces

import (
	"encoding/json"
	"time"

	uuid "github.com/google/uuid" 
)

type UserShift struct {
	ID				uuid.UUID 				`db:"id, primarykey" json:"id"`
	User			*json.RawMessage		`db:"user" json:"user"`
	Shift			*json.RawMessage		`db:"shift" json:"shift"`
	UpdatedAt   	*time.Time       		`db:"updated_at" json:"updated_at"`
	CreatedAt   	time.Time        		`db:"created_at" json:"created_at"`
}