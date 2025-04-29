package interfaces

import (
	uuid "github.com/google/uuid"
)

// LogActivity ...
type LogActivity struct {
	ID       uuid.UUID `db:"id, primarykey" json:"id"`
	UserID   uuid.UUID `db:"user_id" json:"user_id"`
	Name     string    `db:"name" json:"name"`
	Type     string    `db:"type" json:"type"`
	Detail   string    `db:"detail" json:"detail"`
	UrlPhoto string    `db:"url_photo" json:"url_photo"`
}

// ParamsLogActivityAll ...
type ParamsLogActivityAll struct {
	UserID string `json:"user_id"`
}
