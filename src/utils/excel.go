package utils

import (
	"strconv"
	"time"

	model "github.com/Eznopot/RM_api/src/Database/Model"

	logger "github.com/Eznopot/RM_api/src/Logger"
	"github.com/xuri/excelize/v2"
)

func CreateExcelFileAndSaveIt(events []model.CalendarEvent, month int) string {
	f, err := excelize.OpenFile("./Excel/template CRAM.xlsx")
	if err != nil {
		logger.Error(err.Error())
		return ""
	}
	now := time.Now()
	nbDay := time.Date(now.Year(), 11+1, 0, 0, 0, 0, 0, time.UTC).Day()
	for i := 1; i <= nbDay; i++ {
		f.SetCellStr("CRAM_RMS", "A"+strconv.Itoa(i+30), strconv.Itoa(i)+"/"+strconv.Itoa(month)+"/"+strconv.Itoa(now.Year()))
	}

	for i, event := range events {
		date, err := time.Parse("2006-01-02", event.Date)
		if err != nil {
			logger.Error(err.Error())
			return ""
		}
		value, err := strconv.ParseFloat(event.Value, 64)
		if err != nil {
			logger.Error(err.Error())
			return ""
		}
		switch event.EventType {
		case "presence":
			f.SetCellFloat("CRAM_RMS", "B"+strconv.Itoa(date.Day()+30), value, 1, 64)
		case "absence":
			f.SetCellFloat("CRAM_RMS", "C"+strconv.Itoa(date.Day()+30), value, 1, 64)
			f.SetCellValue("CRAM_RMS", "D"+strconv.Itoa(date.Day()+30), event.AbsenceEvent)
		case "prestation supplémentaire":
			f.SetCellFloat("CRAM_RMS", "G"+strconv.Itoa(date.Day()+30), value, 1, 64)
		case "astreinte":
			f.SetCellFloat("CRAM_RMS", "H"+strconv.Itoa(date.Day()+30), value, 1, 64)
		case "autre":
			f.SetCellFloat("CRAM_RMS", "J"+strconv.Itoa(date.Day()+30), value, 1, 64)
			f.SetCellValue("CRAM_RMS", "K"+strconv.Itoa(date.Day()+30), event.OtherEvent)
		}
		f.SetCellValue("CRAM_RMS", "I"+strconv.Itoa(date.Day()+30), event.Comment)
		logger.Debug(strconv.Itoa(i), strconv.Itoa(date.Day()), event.Date, event.Value, event.EventType, event.Comment)
	}
	//ajouter la date de signature et le nom en C23, date en B23

	f.SetCellStr("CRAM_RMS", "B79", strconv.Itoa(now.Day()) + "/" + strconv.Itoa(month) + "/" + strconv.Itoa(now.Year()))

	filename := "./Excel/CRAM" + strconv.Itoa(month) + "-" + strconv.Itoa(now.Year()) + ".xlsx"

	if err := f.SaveAs(filename); err != nil {
		logger.Error(err.Error())
	}
	return filename
}
