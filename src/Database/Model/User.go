package model

type User struct {
	Username string
	Role     string
	Email    string
}

type UserLogin struct {
	Username string
	Email    string
	Role     string
	Token    string
}
