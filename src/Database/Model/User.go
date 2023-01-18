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
	UserInfo  UserInfo
}

type UserInfo struct {
	Address             string
	PostalCode          string
	Country             string
	EmergencyName       string
	EmergencyPhonePerso string
	EmergencyPhonePro   string
	EmergencyLink       string
	EmergencyAddress    string
	EmergencyPostalCode string
	EmergencyCountry    string
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
