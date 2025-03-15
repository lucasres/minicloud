package entities

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

func NewUser(name, email, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}
