package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/hiteshjain48/RSA-aggregator/internal/database"
)

type User struct {
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string `json:"name"`
}

func databaseUserToUser(dbUser database.User) User {
	user := User{
		ID:dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name: dbUser.Name,
	}
	return user
}