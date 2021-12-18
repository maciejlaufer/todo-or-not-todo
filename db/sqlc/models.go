// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type List struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatorID uuid.UUID `json:"creator_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Task struct {
	ID          uuid.UUID      `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	ListID      uuid.UUID      `json:"list_id"`
	Completed   bool           `json:"completed"`
	CreatedAt   time.Time      `json:"created_at"`
}

type User struct {
	ID        uuid.UUID      `json:"id"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	FirstName sql.NullString `json:"first_name"`
	LastName  sql.NullString `json:"last_name"`
	CreatedAt time.Time      `json:"created_at"`
}

type UsersList struct {
	UserID    uuid.UUID `json:"user_id"`
	ListID    uuid.UUID `json:"list_id"`
	CreatedAt time.Time `json:"created_at"`
}
