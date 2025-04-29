package interfaces

import (
	"encoding/json"
	"time"

	uuid "github.com/google/uuid" 
)

type User struct {
	ID			uuid.UUID 				`db:"id, primarykey" json:"id"`
	Name		string					`db:"username" json:"username"`
	Email		string					`db:"email" json:"email"`
	Password    string      			`db:"password" json:"-"`
	Active		string					`db:"active" json:"active"`
	Role        *json.RawMessage      	`db:"role" json:"role"`
	Company		*json.RawMessage		`db:"company" json:"company"`
	Branch      *json.RawMessage      	`db:"branch" json:"branch"`
	Department  *json.RawMessage      	`db:"department" json:"department"`
	Position	*json.RawMessage		`db:"position" json:"position"`
	UpdatedAt   *time.Time       		`db:"updated_at" json:"updated_at"`
	CreatedAt   time.Time        		`db:"created_at" json:"created_at"`
}