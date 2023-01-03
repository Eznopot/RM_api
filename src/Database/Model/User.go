package model

type User struct {
	Id        int
	Username  string
	Role      string
	Email     string
	Phone     string
	Firstname string
	Lastname  string
	HaveCV    bool
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
