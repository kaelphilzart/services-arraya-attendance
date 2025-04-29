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
	ClockInTime 	*time.Time      		`db:"clock_in_time" json:"clock_in_time"`
	ClockOutTime	*time.Time      		`db:"clock_out_time" json:"clock_out_time"`
	UpdatedAt   	*time.Time       		`db:"updated_at" json:"updated_at"`
	CreatedAt   	time.Time        		`db:"created_at" json:"created_at"`
}