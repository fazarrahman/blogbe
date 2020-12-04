package model

// User ...
type User struct {
	ID        string `json:"id"`
	Email     string `json:"email" validate:"required,email"`
	Username  string `json:"username" validate:"required"`
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"omitempty"`
	Password  string `json:"password,omitempty" validate:"required"`
	Status    int8   `json:"status"`
}
