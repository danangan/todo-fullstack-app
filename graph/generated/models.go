// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

type AuthPayload struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
