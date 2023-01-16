package database

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	model "github.com/Eznopot/RM_api/src/Database/Model"
	logger "github.com/Eznopot/RM_api/src/Logger"
	"github.com/Eznopot/RM_api/src/utils"
	"github.com/google/uuid"
)

//* utilitaries functions

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

//* Session functions

func addToken(user_id int64, token string) bool {
	db := GetDb()
	stmt, err := db.Prepare("INSERT INTO Token (user_id, uuid, expiration) VALUES (?, ?, ?)")
	expirationDate := time.Now().Add(time.Hour * 24 * 2)
	if err != nil {
		logger.Error(err.Error())
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
		logger.Error(err.Error())
	}
	res, err := stmt.Exec(time.Now(), token)
	if err != nil {
		logger.Error(err.Error())
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

//* User functions

func Register(id int, username, firstname, lastname, phone, email, password, address, postalCode, country, emergencyName, emergencyNumber, emergencyNumberPro, emergencyLink, addressEmergency, postalCodeEmergency, countryEmergency string) (bool, string) {
	db := GetDb()
	if _, res := rowExists("SELECT * FROM User WHERE username = ?", username); res {
		logger.Info("User already exists")
		return false, "User already exists"
	}
	hashpassword := MD5(password)
	stmt, err := db.Prepare("INSERT INTO User (username, email, password, role, firstname, lastname, phone) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		logger.Error(err.Error())
		return false, "Error"
	}
	res, err := stmt.Exec(username, email, hashpassword, 1, firstname, lastname, phone)
	if err != nil {
		return false, "Error"
	}
	user_id, _ := res.LastInsertId()

	stmt, err = db.Prepare("INSERT INTO UserInformation (user_id, address, country, postal_code, emergency_contact_name, emergency_contact_phone_perso, emergency_contact_phone_pro, emergency_contact_address, emergency_contact_country, emergency_contact_postal_code, emergency_link_family) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		logger.Error(err.Error())
		return false, "Error"
	}
	_, err = stmt.Exec(user_id, address, country, postalCode, emergencyName, emergencyNumber, emergencyNumberPro, addressEmergency, countryEmergency, postalCodeEmergency, emergencyLink)
	if err != nil {
		stmt.Exec("ROLLBACK")
		return false, "Error"
	}
	stmt, err = db.Prepare("UPDATE CV set candidat_id = null, user_id = ? where candidat_id = ?")
	if err != nil {
		logger.Error(err.Error())
		return false, "Error"
	}
	_, err = stmt.Exec(user_id, id)
	if err != nil {
		stmt.Exec("ROLLBACK")
		return false, "Error"
	}
	if id != -1 {
		stmt, err = db.Prepare("DELETE FROM RDV where candidat_id = ?")
		if err != nil {
			logger.Error(err.Error())
			return false, "Error"
		}
		_, err = stmt.Exec(id)
		if err != nil {
			logger.Error(err.Error())
			stmt.Exec("ROLLBACK")
			return false, "Error"
		}
	}
	if id != -1 {
		stmt, err = db.Prepare("DELETE FROM Candidat where id = ?")
		if err != nil {
			logger.Error(err.Error())
			return false, "Error"
		}
		_, err = stmt.Exec(id)
		if err != nil {
			logger.Error(err.Error())
			stmt.Exec("ROLLBACK")
			return false, "Error"
		}
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
	if user_id, res := rowExists("SELECT id FROM User WHERE username = ? or email = ? AND password = ?", username, username, MD5(password)); res {
		var role string
		var email string
		var username string
		var firstname string
		var lastname string
		var phone string
		db := GetDb()
		db.QueryRow("SELECT role+0, email, username, firstname, lastname, phone FROM User WHERE id = ?", user_id).Scan(&role, &email, &username, &firstname, &lastname, &phone)
		token := uuid.New().String()
		addToken(int64(user_id), token)
		return true, model.UserLogin{Email: email, Username: username, Role: role, Token: token, Firstname: firstname, Lastname: lastname, Phone: phone}
	}
	return false, model.UserLogin{}
}

func Logout(token string) (bool, string) {
	db := GetDb()
	stmt, err := db.Prepare("DELETE FROM Token WHERE uuid = ?")
	if err != nil {
		logger.Error(err.Error())
	}
	_, err = stmt.Exec(token)
	if err != nil {
		logger.Error(err.Error())
	}
	return true, "Logout success"
}

func GetInfo(token string) (bool, model.User) {
	db := GetDb()
	user_id, _ := CheckSession(token)
	var user model.User
	db.QueryRow("SELECT username, role, email, firstname, lastname FROM User WHERE id = ?", user_id).Scan(&user.Username, &user.Role, &user.Email, &user.Firstname, &user.Lastname)
	return true, user
}

func GetInfoByEmail(email string) (bool, model.User) {
	db := GetDb()
	var user model.User
	db.QueryRow("SELECT username, role, email, firstname, lastname FROM User WHERE email = ?", email).Scan(&user.Username, &user.Role, &user.Email, &user.Firstname, &user.Lastname)
	return true, user
}

// ! never send the result of this function to the client
func GetAllUserWithAllData() (bool, []model.User) {
	db := GetDb()
	res := []model.User{}
	row, err := db.Query("SELECT id, username, role, email, firstname, lastname, phone FROM User")
	if err != nil {
		return false, res
	}
	for row.Next() {
		var user model.User
		row.Scan(&user.Id, &user.Username, &user.Role, &user.Email, &user.Firstname, &user.Lastname, &user.Phone)
		res = append(res, user)
	}
	return true, res
}

func GetAllUser() (bool, []model.User) {
	db := GetDb()
	res := []model.User{}
	row, err := db.Query("SELECT id, username, role, email, firstname, lastname, phone FROM User")
	if err != nil {
		return false, res
	}
	for row.Next() {
		var user model.User
		row.Scan(&user.Id, &user.Username, &user.Role, &user.Email, &user.Firstname, &user.Lastname, &user.Phone)
		err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM CV WHERE user_id = ?) AS value_exists", user.Id).Scan(&user.HaveCV)
		user.Id = 0
		if err != nil {
			logger.Error(err.Error())
		}
		res = append(res, user)
	}
	return true, res
}

func GetUserCv(email string) (bool, model.Candidat) {
	db := GetDb()
	var CV model.Candidat
	err := db.QueryRow("SELECT CV.created_time, CV.competence, CV.experience, CV.formation FROM CV JOIN User ON CV.user_id = User.id WHERE User.email = ?", email).Scan(&CV.CreatedTime, &CV.Competence, &CV.Experience, &CV.Formation)
	if err != nil {
		return false, CV
	}
	return true, CV
}

func GetRoles() (bool, []string) {
	res := getEnumValue("role", "User")
	return true, res
}

func UpdateRole(email, role string) (bool, string) {
	db := GetDb()
	stmt, err := db.Prepare("UPDATE User SET role = ? WHERE email = ?")
	if err != nil {
		return false, "Role cant be set"
	}
	stmt.Exec(role, email)
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
		res = append(res, "CongesAdmin")
		res = append(res, "Offer")
	}
	if role >= 3 {
		res = append(res, "SaPanelAdmin")
	}

	return true, res
}

//* Admin functions

func AddAdminString(value string) (bool, int64) {
	db := GetDb()
	stmt, err := db.Prepare("INSERT INTO AdminInfo (value) VALUES (?)")
	if err != nil {
		logger.Error(err.Error())
		return false, -1
	}
	res, err := stmt.Exec(value)
	if err != nil {
		logger.Error(err.Error())
		return false, -1
	}
	id, _ := res.LastInsertId()
	return true, id
}

func ModifyAdminString(id int, value string) (bool, string) {
	db := GetDb()
	stmt, err := db.Prepare("UPDATE AdminInfo SET value = ? where id = ?")
	if err != nil {
		logger.Error(err.Error())
		return false, "Error"
	}
	stmt.Exec(value, id)
	return true, "Value successfully modify"
}

func DeleteAdminString(id int) (bool, string) {
	db := GetDb()
	stmt, err := db.Prepare("DELETE FROM AdminInfo WHERE id = ?")
	if err != nil {
		logger.Error(err.Error())
		return false, "Error"
	}
	stmt.Exec(id)
	return true, "Value successfully deleted"
}

//* Offer functions

func AddOffer(token, title, description string, price float64) (bool, int64) {
	db := GetDb()
	_, user_id := CheckSession(token)
	stmt, err := db.Prepare("INSERT INTO Offer (title, description, creator_id, price) VALUES (?, ?, ?, ?)")
	if err != nil {
		logger.Error(err.Error())
		return false, 0
	}
	res, err := stmt.Exec(title, description, user_id, price)
	if err != nil {
		logger.Error(err.Error())
		return false, 0
	}
	id, _ := res.LastInsertId()
	return true, id
}

func ModifyOffer(id int, title, description string, price float64) (bool, string) {
	db := GetDb()
	stmt, err := db.Prepare("UPDATE Offer SET title = ?, description = ?, price = ? WHERE id = ?")
	if err != nil {
		logger.Error(err.Error())
		return false, "Error"
	}
	stmt.Exec(title, description, price, id)
	return true, "Offer successfully modified"
}

func DeleteOffer(id int) (bool, string) {
	db := GetDb()
	stmt, err := db.Prepare("DELETE FROM Offer WHERE id = ?")
	if err != nil {
		logger.Error(err.Error())
		return false, "Error"
	}
	_, err = stmt.Exec(id)
	if err != nil {
		logger.Error(err.Error())
		return false, "Error"
	}
	return true, "Offer successfully deleted"
}

func GetOffers() (bool, []model.Offer) {
	db := GetDb()
	rows, err := db.Query("SELECT id, title, description, created_time, creator_id, price FROM Offer")

	if err != nil {
		logger.Error(err.Error())
		return false, nil
	}
	var res []model.Offer
	for rows.Next() {
		var offer model.Offer
		err := rows.Scan(&offer.Id, &offer.Title, &offer.Description, &offer.CreatedTime, &offer.UserId, &offer.Price)
		if err != nil {
			logger.Error(err.Error())
			return false, nil
		}

		res = append(res, offer)
	}
	return true, res
}

func LoadSomeOffers(limit, offset string) (bool, []model.Offer) {
	db := GetDb()
	var res []model.Offer
	row, err := db.Query("SELECT id, title, description, price, created_time FROM Offer LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		logger.Error(err.Error())
		return false, res
	}
	for row.Next() {
		var tmp = model.Offer{}
		row.Scan(&tmp.Id, &tmp.Title, &tmp.Description, &tmp.Price, &tmp.CreatedTime)
		res = append(res, tmp)
	}
	return true, res
}

//* Candidat functions

func AddCandidat(firstname, lastname, email, phone, formation, experience, competence string, offerId int) (bool, string) {
	db := GetDb()
	if _, res := rowExists("SELECT * FROM Candidat WHERE email = ?", email); res {
		logger.Error("Candidat already exists")
		return false, "Candidat already exists"
	}
	stmt, err := db.Prepare("INSERT INTO Candidat (firstname, lastname, email, phone, offerId) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		logger.Error(err.Error())
		return false, "Error"
	}
	res, err := stmt.Exec(firstname, lastname, email, phone, offerId)
	if err != nil {
		return false, "Error"
	}
	id, _ := res.LastInsertId()
	stmt, err = db.Prepare("INSERT INTO CV (candidat_id, competence, experience, formation) VALUES (?, ?, ?, ?)")
	if err != nil {
		logger.Error(err.Error())
		return false, "Error"
	}
	_, err = stmt.Exec(id, competence, experience, formation)
	if err != nil {
		return false, "Error"
	}
	return true, "Candidat successfully added"
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
		var offer = model.Offer{}
		row.Scan(&tmp.Id, &tmp.Firstname, &tmp.Lastname, &tmp.Email, &tmp.Phone, &offer.Id)
		if offer.Id != 0 {
			db.QueryRow("SELECT id, title, description, price, created_time WHERE id = ?", &offer.Id).Scan(&offer.Id, &offer.Title, &offer.Description, &offer.Price, &offer.CreatedTime)
			tmp.CandidateToOffer = offer
		}
		db.QueryRow("SELECT id, firstname, lastname, email, created_time FROM Candidat WHERE id = ?", tmpId).Scan(&tmp.Id, &elem[0], &elem[1], &tmp.Email, &tmp.CreatedTime)
		tmp.Initial = elem[0][0:1] + elem[1][0:1]
		res = append(res, tmp)
	}
	return true, res
}

func SearchCandidatByEmail(email string) (bool, []model.Candidat) {
	db := GetDb()
	var res []model.Candidat
	email = "%" + email + "%"
	row, err := db.Query("SELECT id, firstname, lastname, email FROM Candidat WHERE email LIKE ?", email)
	if err != nil {
		return false, res
	}
	for row.Next() {
		var elem [2]string
		var tmp = model.Candidat{}
		row.Scan(&tmp.Id, &elem[0], &elem[1], &tmp.Email)
		db.QueryRow("SELECT competence, experience, formation, created_time FROM CV WHERE candidat_id = ?", &tmp.Id).Scan(&tmp.Competence, &tmp.Experience, &tmp.Formation, &tmp.CreatedTime)
		tmp.Initial = elem[0][0:1] + elem[1][0:1]
		res = append(res, tmp)
	}
	return true, res
}

func LoadSomeCandidat(limit, offset string) (bool, []model.Candidat) {
	db := GetDb()
	var res []model.Candidat
	row, err := db.Query("SELECT id, firstname, lastname, email, phone, offerId FROM Candidat LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return false, res
	}
	for row.Next() {
		var tmp = model.Candidat{}
		var offer = model.Offer{}
		row.Scan(&tmp.Id, &tmp.Firstname, &tmp.Lastname, &tmp.Email, &tmp.Phone, &offer.Id)
		if offer.Id != 0 {
			db.QueryRow("SELECT id, title, description, price, created_time WHERE id = ?", &offer.Id).Scan(&offer.Id, &offer.Title, &offer.Description, &offer.Price, &offer.CreatedTime)
			tmp.CandidateToOffer = offer
		}
		db.QueryRow("SELECT competence, experience, formation, created_time FROM CV WHERE candidat_id = ?", &tmp.Id).Scan(&tmp.Competence, &tmp.Experience, &tmp.Formation, &tmp.CreatedTime)
		tmp.Initial = tmp.Firstname[0:1] + tmp.Lastname[0:1]
		res = append(res, tmp)
	}
	return true, res
}

//* Calendar functions

func GetCalendarEventsByUserId(user_id int, month int) (bool, []model.CalendarEvent) {
	db := GetDb()
	actualYear := time.Now().Year()
	row, err := db.Query("SELECT id, date, event_type, comment, value, other_event, consultant_backup, absence_event FROM Calendar WHERE user_id = ? AND YEAR(date) = ? AND MONTH(date) = ? ", user_id, actualYear, month)
	if err != nil {
		logger.Error(err.Error())
		return false, []model.CalendarEvent{}
	}
	res := []model.CalendarEvent{}
	for row.Next() {
		tmp := model.CalendarEvent{}
		var otherEvent, consultantBackup, absenceEvent sql.NullString
		row.Scan(&tmp.Id, &tmp.Date, &tmp.EventType, &tmp.Comment, &tmp.Value, &otherEvent, &consultantBackup, &absenceEvent)
		tmp.OtherEvent = otherEvent.String
		tmp.ConsultantBackup = consultantBackup.String
		tmp.AbsenceEvent = absenceEvent.String
		res = append(res, tmp)
	}
	return true, res
}

func GetCalendarEvents(token string, month int) (bool, []model.CalendarEvent) {
	db := GetDb()
	_, user_id := CheckSession(token)
	actualYear := time.Now().Year()
	row, err := db.Query("SELECT id, date, event_type, comment, value, other_event, consultant_backup, absence_event FROM Calendar WHERE user_id = ? AND YEAR(date) = ? AND MONTH(date) = ? ", user_id, actualYear, month)
	if err != nil {
		logger.Error(err.Error())
		return false, []model.CalendarEvent{}
	}
	res := []model.CalendarEvent{}
	for row.Next() {
		tmp := model.CalendarEvent{}
		var otherEvent, consultantBackup, absenceEvent sql.NullString
		row.Scan(&tmp.Id, &tmp.Date, &tmp.EventType, &tmp.Comment, &tmp.Value, &otherEvent, &consultantBackup, &absenceEvent)
		tmp.OtherEvent = otherEvent.String
		tmp.ConsultantBackup = consultantBackup.String
		tmp.AbsenceEvent = absenceEvent.String
		res = append(res, tmp)
	}
	return true, res
}

func GetCalendarEventsByEmail(email string, month int) (bool, []model.CalendarEvent) {
	db := GetDb()
	actualYear := time.Now().Year()
	row, err := db.Query("SELECT Calendar.id, Calendar.date, Calendar.event_type, Calendar.comment, Calendar.value, Calendar.other_event, Calendar.consultant_backup, Calendar.absence_event FROM Calendar JOIN User ON Calendar.user_id = User.id where User.email = ? AND YEAR(Calendar.date) = ? AND MONTH(Calendar.date) = ?", email, actualYear, month)
	if err != nil {
		logger.Error(err.Error())
		return false, []model.CalendarEvent{}
	}
	res := []model.CalendarEvent{}
	for row.Next() {
		tmp := model.CalendarEvent{}
		var otherEvent, consultantBackup, absenceEvent sql.NullString
		row.Scan(&tmp.Id, &tmp.Date, &tmp.EventType, &tmp.Comment, &tmp.Value, &otherEvent, &consultantBackup, &absenceEvent)
		tmp.OtherEvent = otherEvent.String
		tmp.ConsultantBackup = consultantBackup.String
		tmp.AbsenceEvent = absenceEvent.String
		res = append(res, tmp)
	}
	return true, res
}

func AddCalendarEvent(token string, date string, eventType string, comment string, value float64, otherEvent interface{}, backupName interface{}, absenceType interface{}) (bool, int64) {
	db := GetDb()
	_, user_id := CheckSession(token)
	date = strings.ReplaceAll(date, "Z", "")
	stmt, err := db.Prepare("INSERT INTO Calendar (user_id, date, event_type, comment, value, consultant_backup, other_event, absence_event) VALUES (?, ?, ?, ?, ?, ?, CASE WHEN ? != '' THEN ? ELSE NULL END, CASE WHEN ? != '' THEN ? ELSE NULL END)")
	if err != nil {
		logger.Error(err.Error())
		return false, -1
	}
	res, err := stmt.Exec(user_id, date, eventType, comment, value, backupName, otherEvent.(string), otherEvent.(string), absenceType.(string), absenceType.(string))
	if err != nil {
		logger.Error(err.Error())
		return false, -1
	}
	lastInsert, _ := res.LastInsertId()
	return true, lastInsert
}

func AutoPresenceCalendarEvents(token string, month int) (bool, []model.CalendarEvent) {
	db := GetDb()
	_, user_id := CheckSession(token)
	stmt, err := db.Prepare("DELETE FROM Calendar WHERE user_id = ? AND event_type = ? AND YEAR(date) = ? AND MONTH(date) = ?")
	if err != nil {
		logger.Error(err.Error())
		return false, []model.CalendarEvent{}
	}
	_, err = stmt.Exec(user_id, "presence", time.Now().Year(), time.Month(month))
	if err != nil {
		logger.Error(err.Error())
		return false, []model.CalendarEvent{}
	}
	stmt, err = db.Prepare("INSERT INTO Calendar (user_id, date, event_type, comment, value, consultant_backup) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		logger.Error(err.Error())
		return false, []model.CalendarEvent{}
	}
	start := time.Date(time.Now().Year(), time.Month(month), 1, 0, 0, 0, 0, &time.Location{})
	end := start.AddDate(0, 1, -1)
	dayOff := utils.GetHollidays(start.Year())
	if dayOff == nil {
		logger.Error("Bad year givent for getting holliday")
		return false, []model.CalendarEvent{}
	}
	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		if d.Weekday() != time.Saturday && d.Weekday() != time.Sunday && !utils.IsDayOff(dayOff, d) {
			_, err = stmt.Exec(user_id, d, "presence", "", 1, "")
			if err != nil {
				logger.Error(err.Error())
				return false, []model.CalendarEvent{}
			}
		}
	}
	_, events := GetCalendarEvents(token, month)
	return true, events
}

func DeleteCalendarEvent(token string, id int) (bool, string) {
	db := GetDb()
	_, user_id := CheckSession(token)
	stmt, err := db.Prepare("DELETE FROM Calendar WHERE id = ? AND user_id = ?")
	if err != nil {
		logger.Error(err.Error())
		return false, "Error"
	}
	_, err = stmt.Exec(id, user_id)
	if err != nil {
		logger.Error(err.Error())
		return false, "Error"
	}
	return true, "Event successfully deleted"
}

func ModifyCalendarEvent(token string, id int, date string, eventType string, comment string, value float64, otherEvent interface{}, backupName interface{}, absenceType interface{}) (bool, string) {
	db := GetDb()
	_, user_id := CheckSession(token)
	date = strings.ReplaceAll(date, "Z", "")
	stmt, err := db.Prepare("UPDATE Calendar SET date = ?, event_type = ?, comment = ?, value = ?, consultant_backup = ?, other_event = CASE WHEN ? != '' THEN ? ELSE NULL END, absence_event = CASE WHEN ? != '' THEN ? ELSE NULL END WHERE id = ? AND user_id = ?")
	if err != nil {
		logger.Error(err.Error())
		return false, "Error"
	}
	_, err = stmt.Exec(date, eventType, comment, value, backupName, otherEvent.(string), otherEvent.(string), absenceType.(string), absenceType.(string), id, user_id)
	if err != nil {
		logger.Error(err.Error())
		return false, "Error"
	}

	return true, "Event successfully modified"
}

// * Calendar Enum functions
func getEnumValue(enumName, tableName string) []string {
	db := GetDb()
	var row string
	db.QueryRow("SELECT TRIM(TRAILING ')' FROM TRIM(LEADING '(' FROM TRIM(LEADING 'enum' FROM column_type))) column_type FROM	information_schema.columns WHERE table_name = ? AND column_name = ?;", tableName, enumName).Scan(&row)
	return strings.Split(strings.ReplaceAll(row, "'", ""), ",")
}

func GetEventTypes() (bool, []string) {
	res := getEnumValue("event_type", "Calendar")
	return true, res
}

func GetOtherEventTypes() (bool, []string) {
	res := getEnumValue("other_event", "Calendar")
	return true, res
}

func GetAbsenceEventTypes() (bool, []string) {
	res := getEnumValue("absence_event", "Calendar")
	return true, res
}

//* HolliDay functions

func AddHollidayRequest(token string, dateStart string, dateEnd string) (bool, int64) {
	db := GetDb()
	_, user_id := CheckSession(token)
	dateStart = strings.ReplaceAll(dateStart, "Z", "")
	dateEnd = strings.ReplaceAll(dateEnd, "Z", "")
	stmt, err := db.Prepare("INSERT INTO Holliday (user_id, dateStart, dateEnd) VALUES (?, ?, ?)")
	if err != nil {
		logger.Error(err.Error())
		return false, -1
	}
	res, err := stmt.Exec(user_id, dateStart, dateEnd)
	if err != nil {
		logger.Error(err.Error())
		return false, -1
	}
	lastInsert, _ := res.LastInsertId()
	return true, lastInsert
}

func AcceptHollidayRequest(id int) (bool, string) {
	db := GetDb()
	stmt, err := db.Prepare("UPDATE Holliday SET status = 'accepted' WHERE id = ?")
	if err != nil {
		logger.Error(err.Error())
		return false, "Error"
	}
	_, err = stmt.Exec(id)
	if err != nil {
		logger.Error(err.Error())
		return false, "Error"
	}
	return true, "Request successfully accepted"
}

func DeclineHollidayRequest(id int) (bool, string) {
	db := GetDb()
	stmt, err := db.Prepare("UPDATE Holliday SET status = 'refused' WHERE id = ?")
	if err != nil {
		logger.Error(err.Error())
		return false, "Error"
	}
	_, err = stmt.Exec(id)
	if err != nil {
		logger.Error(err.Error())
		return false, "Error"
	}
	return true, "Request successfully declined"
}

func ModifyHollidayRequest(token string, id int, dateStart string, dateEnd string) (bool, string) {
	db := GetDb()
	_, user_id := CheckSession(token)
	dateStart = strings.ReplaceAll(dateStart, "Z", "")
	dateEnd = strings.ReplaceAll(dateEnd, "Z", "")
	stmt, err := db.Prepare("UPDATE Holliday SET dateStart = ?, dateEnd = ? WHERE id = ? AND user_id = ? AND status = 'pending'")
	if err != nil {
		logger.Error(err.Error())
		return false, "Error"
	}
	_, err = stmt.Exec(dateStart, dateEnd, id, user_id)
	if err != nil {
		logger.Error(err.Error())
		return false, "Error"
	}
	return true, "Request successfully modified"
}

func DeleteHollidayRequest(token string, id int) (bool, string) {
	db := GetDb()
	_, user_id := CheckSession(token)
	stmt, err := db.Prepare("DELETE FROM Holliday WHERE id = ? AND user_id = ?")
	if err != nil {
		logger.Error(err.Error())
		return false, "Error"
	}
	_, err = stmt.Exec(id, user_id)
	if err != nil {
		logger.Error(err.Error())
		return false, "Error"
	}
	return true, "Request successfully deleted"
}

func DeleteOtherHollidayRequest(id int) (bool, string) {
	db := GetDb()
	stmt, err := db.Prepare("DELETE FROM Holliday WHERE id = ?")
	if err != nil {
		logger.Error(err.Error())
		return false, "Error"
	}
	_, err = stmt.Exec(id)
	if err != nil {
		logger.Error(err.Error())
		return false, "Error"
	}
	return true, "Request successfully deleted"
}

func GetAllHollidayRequest(month int) (bool, []model.HollidayRequest) {
	db := GetDb()
	var res []model.HollidayRequest
	rows, err := db.Query("SELECT Holliday.id, dateStart, dateEnd, status+0, username, email FROM Holliday LEFT JOIN `User` ON `User`.id = Holliday.user_id WHERE MONTH(dateStart) >= ?", month)
	if err != nil {
		logger.Error(err.Error())
		return false, nil
	}
	for rows.Next() {
		var tmp model.HollidayRequest
		rows.Scan(&tmp.Id, &tmp.DateStart, &tmp.DateEnd, &tmp.Status, &tmp.Username, &tmp.Email)
		res = append(res, tmp)
	}
	return true, res
}

func GetHollidayRequest(token string) (bool, []model.HollidayRequest) {
	db := GetDb()
	_, user_id := CheckSession(token)
	var res []model.HollidayRequest
	rows, err := db.Query("SELECT id, dateStart, dateEnd, status+0 FROM Holliday WHERE user_id = ?", user_id)
	if err != nil {
		logger.Error(err.Error())
		return false, nil
	}
	for rows.Next() {
		var tmp model.HollidayRequest
		rows.Scan(&tmp.Id, &tmp.DateStart, &tmp.DateEnd, &tmp.Status)
		res = append(res, tmp)
	}
	return true, res
}

// * RDV function

func GetRDVEvent(month int) (bool, []model.RDVEvent) {
	db := GetDb()
	var res []model.RDVEvent
	rows, err := db.Query("SELECT id, user_id, candidat_id, date, appreciation FROM RDV WHERE MONTH(date) >= ?", month)
	if err != nil {
		logger.Error(err.Error())
		return false, nil
	}
	for rows.Next() {
		var tmp model.RDVEvent
		var candidat int
		rows.Scan(&tmp.Id, &tmp.UserId, &candidat, &tmp.Date, &tmp.Appreciation)
		db.QueryRow("SELECT email, firstname, lastname FROM Candidat WHERE id = ?", candidat).Scan(&tmp.Candidat.Email, &tmp.Candidat.Firstname, &tmp.Candidat.Lastname)
		db.QueryRow("SELECT competence, formation, experience FROM CV WHERE candidat_id = ?", candidat).Scan(&tmp.Candidat.Competence, &tmp.Candidat.Formation, &tmp.Candidat.Experience)
		res = append(res, tmp)
	}
	return true, res
}

func AddRDVEvent(token string, candidat_id int, date string) (bool, int64) {
	db := GetDb()
	_, user_id := CheckSession(token)
	date = strings.ReplaceAll(date, "Z", "")
	stmt, err := db.Prepare("INSERT INTO RDV (user_id, candidat_id, date) VALUES (?, ?, ?)")
	if err != nil {
		logger.Error(err.Error())
		return false, -1
	}
	res, err := stmt.Exec(user_id, candidat_id, date)
	if err != nil {
		logger.Error(err.Error())
		return false, -1
	}
	lastInsert, _ := res.LastInsertId()
	return true, lastInsert
}

func InsertAppreciation(token string, id int, appreciation string) (bool, string) {
	db := GetDb()
	_, user_id := CheckSession(token)
	stmt, err := db.Prepare("UPDATE RDV SET appreciation = ? WHERE id = ? AND user_id = ?")
	if err != nil {
		logger.Error(err.Error())
		return false, "Error"
	}
	_, err = stmt.Exec(appreciation, id, user_id)
	if err != nil {
		logger.Error(err.Error())
		return false, "Error"
	}
	return true, "Appreciation successfully added"
}

func ModifyRDVEvent(token string, id int, date string) (bool, string) {
	db := GetDb()
	_, user_id := CheckSession(token)
	date = strings.ReplaceAll(date, "Z", "")
	stmt, err := db.Prepare("UPDATE RDV SET date = ? WHERE id = ? AND user_id = ?")
	if err != nil {
		logger.Error(err.Error())
		return false, "Error"
	}
	_, err = stmt.Exec(date, id, user_id)
	if err != nil {
		logger.Error(err.Error())
		return false, "Error"
	}
	return true, "RDV successfully modified"
}

func DeleteRDVEvent(token string, id int) (bool, string) {
	db := GetDb()
	_, user_id := CheckSession(token)
	stmt, err := db.Prepare("DELETE FROM RDV WHERE id = ? AND user_id = ?")
	if err != nil {
		logger.Error(err.Error())
		return false, "Error"
	}
	_, err = stmt.Exec(id, user_id)
	if err != nil {
		logger.Error(err.Error())
		return false, "Error"
	}
	return true, "RDV successfully deleted"
}

//*Other function

func GetAdminString() (bool, []model.AdminInfo) {
	db := GetDb()
	var adminInfoList []model.AdminInfo
	rows, err := db.Query("SELECT id, value FROM AdminInfo")
	if err != nil {
		logger.Error(err.Error())
		return false, adminInfoList
	}
	for rows.Next() {
		var tmp = model.AdminInfo{}
		rows.Scan(&tmp.Id, &tmp.Info)
		adminInfoList = append(adminInfoList, tmp)
	}
	return true, adminInfoList
}
