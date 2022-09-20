package database

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"log"

	model "github.com/Eznopot/RM_api/src/Database/Model"
	"github.com/google/uuid"
)

func rowExists(query string, args ...interface{}) (int, bool) {
	var exists bool
	var result int
	db := GetDb()
	newquery := fmt.Sprintf("SELECT exists (%s)", query)
	err := db.QueryRow(newquery, args...).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("error checking if row exists '%s' %v", args, err)
	}
	db.QueryRow(query, args...).Scan(&result)
	return result, exists
}

func MD5(text string) string {
	data := []byte(text)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func addToken(user_id int64, token string) bool {
	db := GetDb()
	stmt, err := db.Prepare("INSERT INTO Token (user_id, uuid) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
		return false
	}
	_, err = stmt.Exec(user_id, token)
	if err != nil {
		return false
	}
	if err != nil {
		return false
	}
	return true
}

func CheckSession(token string) (int, bool) {
	id, res := rowExists("SELECT user_id FROM Token WHERE uuid = ?", token)
	return id, res
}

func CheckRightIsAdmin(user_id int) (int, bool) {
	role, res := rowExists("SELECT role+0 FROM User WHERE id = ?", user_id)
	return role, res
}

func Register(username, password string) (bool, string) {
	db := GetDb()
	if _, res := rowExists("SELECT * FROM User WHERE username = ?", username); res {
		log.Println("User already exists")
		return false, "User already exists"
	}
	hashpassword := MD5(password)
	token := uuid.New().String()
	stmt, err := db.Prepare("INSERT INTO User (username, password, role) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return false, "Error"
	}
	res, err := stmt.Exec(username, hashpassword, 1)
	if err != nil {
		return false, "Error"
	}
	if err != nil {
		return false, "Error"
	}
	id, _ := res.LastInsertId()
	addToken(id, token)
	return true, token
}

func Login(username, password string) (bool, string) {
	if user_id, res := rowExists("SELECT id FROM User WHERE username = ? AND password = ?", username, MD5(password)); res {
		token := uuid.New().String()
		addToken(int64(user_id), token)
		return true, token
	}
	return false, "Bad credential"
}

func Logout(token string) (bool, string) {
	db := GetDb()
	stmt, err := db.Prepare("DELETE FROM Token WHERE uuid = ?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(token)
	if err != nil {
		log.Fatal(err)
	}
	return true, "Logout success"
}

func GetInfo(token string) (bool, model.User) {
	db := GetDb()
	user_id, _ := CheckSession(token)
	var user model.User
	db.QueryRow("SELECT username, role FROM User WHERE id = ?", user_id).Scan(&user.Username, &user.Role)
	return true, user
}

func GetAllUser() (bool, []model.User) {
	db := GetDb()
	res := []model.User{}
	row, err := db.Query("SELECT username, role FROM User")
	if err != nil {
		return false, res
	}
	for row.Next() {
		var user model.User
		row.Scan(&user.Username, &user.Role)
		res = append(res, user)
	}
	return true, res
}

func UpdateRole(username string, role int) (bool, string) {
	db := GetDb()
	stmt, err := db.Prepare("UPDATE User SET role = ? WHERE username = ?")
	if err != nil {
		return false, "Role cant be set"
	}
	stmt.Exec(role, username)
	return true, "Role successfully set"
}

func GetPages() (bool, string) {
	return true, "elem"
}
