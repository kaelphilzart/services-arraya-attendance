package interfaces

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type LeaveApproval struct {
	ID                        uuid.UUID        `db:"id, primarykey" json:"id"`
	Leave                     *json.RawMessage `db:"leave" json:"leave"`                                                
	Level                 	  int 			   `db:"level" json:"level"`
	ApprovedBy                *json.RawMessage `db:"approved_by" json:"approved_by"`                                                        
	Status   			  	  string           `db:"status" json:"status"`
	Note	  				  string       	   `db:"note" json:"note"`
	ApprovedAt   			  *time.Time       `db:"approved_at" json:"approved_at"`
	CreatedAt   			  time.Time        `db:"created_at" json:"created_at"`
	UpdatedAt   			  *time.Time       `db:"updated_at" json:"updated_at"`
}
