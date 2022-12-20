package model

type User struct {
	Username  string
	Role      string
	Email     string
	Phone     string
	Firstname string
	Lastname  string
}

type UserLogin struct {
	Username  string
	Email     string
	Role      string
	Token     string
	Phone     string
	Firstname string
	Lastname  string
}
