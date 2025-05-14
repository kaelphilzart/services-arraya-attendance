package interfaces

import (
	"encoding/json"
	"time"

	uuid "github.com/google/uuid" 
)

type Shift struct {
	ID				uuid.UUID 				`db:"id, primarykey" json:"id"`
	Company			*json.RawMessage		`db:"company" json:"company"`
	Name			string					`db:"name" json:"name"`
	StartTime 		*time.Time      		`db:"start_time" json:"start_time"`
	EndTime			*time.Time      		`db:"end_time" json:"end_time"`
	UpdatedAt   	*time.Time       		`db:"updated_at" json:"updated_at"`
	CreatedAt   	time.Time        		`db:"created_at" json:"created_at"`
}