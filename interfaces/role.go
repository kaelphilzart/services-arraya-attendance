package interfaces

import (
	"time"

	uuid "github.com/google/uuid" 
)

type Role struct {
	ID			uuid.UUID 				`db:"id, primarykey" json:"id"`
	Name		string					`db:"name" json:"name"`
	SlugName   	string      			`db:"slug_name" json:"slug_name"`
	UpdatedAt   *time.Time       		`db:"updated_at" json:"updated_at"`
	CreatedAt   time.Time        		`db:"created_at" json:"created_at"`
}