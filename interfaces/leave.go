package interfaces

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Leave struct {
	ID                        uuid.UUID        `db:"id, primarykey" json:"id"`
	User                      *json.RawMessage `db:"user" json:"user"`                                                
	TypeLeave                 *json.RawMessage `db:"type_leave" json:"type_leave"`                                                
	UrlPhoto                  string      	   `db:"url_photo" json:"url_photo"`                            
	StartDate                 *time.Time       `db:"start_date" json:"start_date"`
	EndDate                   *time.Time       `db:"end_date" json:"end_date"`
	Status   			  	  string           `db:"status" json:"status"`
	CurrentApprovalLevel	  int       	   `db:"current_approval_level" json:"current_approval_level"`
	Description   			  string           `db:"description" json:"description"`
	UpdatedAt   			  *time.Time       `db:"updated_at" json:"updated_at"`
	CreatedAt   			  time.Time        `db:"created_at" json:"created_at"`
}
