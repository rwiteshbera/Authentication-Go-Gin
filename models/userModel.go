package models

import "time"

type User struct {
	UserId    string    `json:"user_id" validate:"required"`
	FirstName string    `json:"first_name" validate:"required, min=2, max=100"`
	LastName  string    `json:"last_name" validate:"required, min=2, max=100"`
	Email     string    `json:"email" validate:"email, required"`
	Password  string    `json:"password" validate:"required, min=6"`
	CreatedAt time.Time `json:"created_at"`
}
