package interfaces

import (
	"time"

	"github.com/google/uuid"
)

type TypeLeave struct {
	ID                        uuid.UUID        `db:"id, primarykey" json:"id"`
	Code                  	  string      	   `db:"code" json:"code"`                            
	Name   			  	  	  string           `db:"name" json:"name"`
	UpdatedAt   			  *time.Time       `db:"updated_at" json:"updated_at"`
	CreatedAt   			  time.Time        `db:"created_at" json:"created_at"`
}
