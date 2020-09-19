package model

// User initializes new user.
type User struct {
	Name     string
	Password string
	Role     string
}

// NewUser creates new user.
func NewUser(name, password, role string) *User {
	return &User{
		Name:     name,
		Password: password,
		Role:     role,
	}
}
