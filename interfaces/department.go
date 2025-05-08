package interfaces

import (
	"encoding/json"
	"time"

	uuid "github.com/google/uuid" 
)

type Department struct {
	ID			uuid.UUID 				`db:"id, primarykey" json:"id"`
	Name		string					`db:"name" json:"name"`
	Company		*json.RawMessage		`db:"company" json:"company"`
	Branch		*json.RawMessage		`db:"branch" json:"branch"`
	Director    *json.RawMessage      	`db:"director" json:"director"`
	UpdatedAt   *time.Time       		`db:"updated_at" json:"updated_at"`
	CreatedAt   time.Time        		`db:"created_at" json:"created_at"`
}