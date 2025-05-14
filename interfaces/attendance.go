package interfaces

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// Absent Status ...
type Attendance struct {
	ID                        uuid.UUID        `db:"id, primarykey" json:"id"`
	User                      *json.RawMessage `db:"user" json:"user"`                                                
	Date                      time.Time        `db:"date" json:"date"`                                                
	ChekInTime                *time.Time       `db:"chek_in_time" json:"chek_in_time"`                            
	ChekOutTime               *time.Time       `db:"chek_out_time" json:"chek_out_time"`
	LatitudeIn   			  string           `db:"latitude_in" json:"latitude_in"`
	LongitudeIn				  string           `db:"longitude_in" json:"longitude_in"`
	LatitudeOut   			  string           `db:"latitude_out" json:"latitude_out"`
	LongitudeOut			  string           `db:"longitude_out" json:"longitude_out"`
	PhotoIn 		          *string          `db:"photo_in" json:"photo_in"`
	PhotoOut        		  *string          `db:"photo_out" json:"photo_out"`
	Note                      string           `db:"note" json:"note"`   
	CreatedAt 				  time.Time        `db:"created_at" json:"created_at"`
	UpdatedAt 				  time.Time        `db:"updated_at" json:"updated_at"`
}
