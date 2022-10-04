package database

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

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
	stmt, err := db.Prepare("INSERT INTO Token (user_id, uuid, expiration) VALUES (?, ?, ?)")
	expirationDate := time.Now().Add(time.Hour * 24 * 2)
	if err != nil {
		log.Fatal(err)
		return false
	}
	_, err = stmt.Exec(user_id, token, expirationDate)
	if err != nil {
		return false
	}
	if err != nil {
		return false
	}
	return true
}

func RemoveExpiredToken(token string) (bool, string) {
	db := GetDb()
	stmt, err := db.Prepare("DELETE FROM Token WHERE expiration < ? && uuid = ?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(time.Now(), token)
	if err != nil {
		log.Fatal(err)
	}
	if nbRow, _ := res.RowsAffected(); nbRow == 0 {
		return true, "Token is valid"
	}
	return false, "Expired token removed"
}

func CheckSession(token string) (bool, int) {
	id, res := rowExists("SELECT user_id FROM Token WHERE uuid = ?", token)
	return res, id
}

func CheckRightIsAdmin(user_id int) (bool, int) {
	role, res := rowExists("SELECT role+0 FROM User WHERE id = ?", user_id)
	return res, role
}

func Register(username, email, password string) (bool, string) {
	db := GetDb()
	if _, res := rowExists("SELECT * FROM User WHERE username = ?", username); res {
		log.Println("User already exists")
		return false, "User already exists"
	}
	hashpassword := MD5(password)
	stmt, err := db.Prepare("INSERT INTO User (username, email, password, role) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return false, "Error"
	}
	_, err = stmt.Exec(username, email, hashpassword, 1)
	if err != nil {
		return false, "Error"
	}
	return true, "User registered"
}

func ChangePassword(token, oldpassword, newpassword string) (bool, string) {
	if user_id, res := rowExists("SELECT user_id FROM Token WHERE uuid = ?", token); res {
		if _, res := rowExists("SELECT id FROM User WHERE id = ? AND password = ?", user_id, MD5(oldpassword)); res {
			db := GetDb()
			stmt, err := db.Prepare("UPDATE User SET password = ? WHERE id = ?")
			if err != nil {
				return false, "Password cant be changed"
			}
			stmt.Exec(MD5(newpassword), user_id)
			token := uuid.New().String()
			addToken(int64(user_id), token)
			return true, token
		}
		return false, "Old password is incorrect"
	}
	return false, "Token is incorrect"
}

func Login(username, password string) (bool, model.UserLogin) {
	if user_id, res := rowExists("SELECT id FROM User WHERE username = ? AND password = ?", username, MD5(password)); res {
		var role string
		var email string
		db := GetDb()
		db.QueryRow("SELECT role+0, email FROM User WHERE id = ?", user_id).Scan(&role, &email)
		token := uuid.New().String()
		addToken(int64(user_id), token)
		return true, model.UserLogin{Email: email, Username: username, Role: role, Token: token}
	}
	return false, model.UserLogin{}
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
	db.QueryRow("SELECT username, role, email FROM User WHERE id = ?", user_id).Scan(&user.Username, &user.Role, &user.Email)
	return true, user
}

func GetAllUser() (bool, []model.User) {
	db := GetDb()
	res := []model.User{}
	row, err := db.Query("SELECT username, role, email FROM User")
	if err != nil {
		return false, res
	}
	for row.Next() {
		var user model.User
		row.Scan(&user.Username, &user.Role, &user.Email)
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

func GetPages(token string) (bool, []string) {
	db := GetDb()
	var role int
	var res []string
	db.QueryRow("SELECT role+0 FROM User WHERE id = (SELECT user_id FROM Token WHERE uuid = ?)", token).Scan(&role)
	if role >= 1 {
		res = append(res, "Calendar")
		res = append(res, "Conges")
		res = append(res, "Setting")
	}
	if role >= 2 {
		res = append(res, "Candidat")
		res = append(res, "RDV")
	}
	if role >= 3 {
		res = append(res, "SaPanelAdmin")
	}
	return true, res
}

func AddCandidat(firstname, lastname, email, formation, experience, competence string) (bool, string, int) {
	db := GetDb()
	if _, res := rowExists("SELECT * FROM Candidat WHERE email = ?", email); res {
		log.Println("Candidat already exists")
		return false, "Candidat already exists", 0
	}
	stmt, err := db.Prepare("INSERT INTO Candidat (firstname, lastname, email) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return false, "Error", 0
	}
	res, err := stmt.Exec(firstname, lastname, email)
	if err != nil {
		return false, "Error", 0
	}
	id, _ := res.LastInsertId()
	stmt, err = db.Prepare("INSERT INTO CV (candidat_id, competence, experience, formation) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return false, "Error", 0
	}
	_, err = stmt.Exec(id, formation, experience, competence)
	if err != nil {
		return false, "Error", 0
	}
	return true, "Candidat successfully added", int(id)
}

func SearchCandidat(search string) (bool, []model.Candidat) {
	db := GetDb()
	var res []model.Candidat
	search = "%" + search + "%"
	row, err := db.Query("SELECT candidat_id, competence, experience, formation FROM CV WHERE competence LIKE ? or experience LIKE ? or formation LIKE ?", search, search, search)
	if err != nil {
		return false, res
	}
	for row.Next() {
		var tmpId int
		var elem [2]string
		var tmp = model.Candidat{}
		row.Scan(&tmpId, &tmp.Competence, &tmp.Experience, &tmp.Formation)
		db.QueryRow("SELECT id, firstname, lastname, email FROM Candidat WHERE id = ?", tmpId).Scan(&tmp.Id, &elem[0], &elem[1], &tmp.Email)
		tmp.Initial = elem[0][0:1] + elem[1][0:1]
		res = append(res, tmp)
	}
	return true, res
}

func AddCalendarEvent(token string, date string, eventType string, comment string, value float64, otherEvent interface{}, backupName interface{}, absenceType interface{}) (bool, string) {
	db := GetDb()
	user_id, _ := CheckSession(token)
	stmt, err := db.Prepare("INSERT INTO Calendar (user_id, date, type, comment, value, other, consultant_backup) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return false, "Error when adding event"
	}
	_, err = stmt.Exec(user_id, date, eventType, comment, value, otherEvent, backupName, absenceType)
	if err != nil {
		log.Fatal(err)
		return false, "Error when adding event"
	}
	return true, "Event successfully added"
}

func getEnumValue(enumName string) ([]string) {
	db := GetDb()
	var row string
	db.QueryRow("SELECT TRIM(TRAILING ')' FROM TRIM(LEADING '(' FROM TRIM(LEADING 'enum' FROM column_type))) column_type FROM	information_schema.columns WHERE table_name = 'Calendar' AND column_name = ?;", enumName).Scan(&row)
	return strings.Split(strings.ReplaceAll(row, "'", ""), ",")
}

func GetEventTypes() (bool, []string) {
	res := getEnumValue("event_type")
	return true, res
}

func GetOtherEventTypes() (bool, []string) {
	res := getEnumValue("other_event")
	return true, res
}


func GetAbsenceEventTypes() (bool, []string) {
	res := getEnumValue("absence_event")
	return true, res
}
