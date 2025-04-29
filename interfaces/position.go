package interfaces

import (
	"encoding/json"
	"time"

	uuid "github.com/google/uuid" 
)

type Position struct {
	ID			uuid.UUID 				`db:"id, primarykey" json:"id"`
	Name		string					`db:"name" json:"name"`
	Department	*json.RawMessage		`db:"department" json:"department"`
	Level   	string      			`db:"level" json:"level"`
	UpdatedAt   *time.Time       		`db:"updated_at" json:"updated_at"`
	CreatedAt   time.Time        		`db:"created_at" json:"created_at"`
}
