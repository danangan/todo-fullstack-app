// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"time"
)

type AuthPayload struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

type Todo struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"dueDate"`
}

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
