package model

type CalendarEvent struct {
	Id               int
	Date             string
	EventType        string
	Comment          string
	Value            string
	OtherEvent       string
	ConsultantBackup string
	AbsenceEvent     string
}
