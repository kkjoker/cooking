package model

import "time"

// User represents a row from 'cooking.users'.
type User struct {
	ID        int       `json:"id"` // id
	Name      string    `json:"name"` // name
	Address   string    `json:"address"` // address
	CreatedAt *time.Time `json:"created_at"` // created_at
	UpdatedAt *time.Time `json:"updated_at"` // updated_at
	DeletedAt *time.Time `json:"deleted_at"` // deleted_at
}
