package interfaces

import (
	"encoding/json"
	"time"

	uuid "github.com/google/uuid" 
)

type Branch struct {
	ID			uuid.UUID 				`db:"id, primarykey" json:"id"`
	Company		*json.RawMessage		`db:"company" json:"company"`
	Name		string					`db:"name" json:"name"`
	Address     string      			`db:"address" json:"address"`
	Contact 	string					`db:"contact" json:"contact"`
	UpdatedAt   *time.Time       		`db:"updated_at" json:"updated_at"`
	CreatedAt   time.Time        		`db:"created_at" json:"created_at"`
}