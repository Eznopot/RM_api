package utils

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	logger "github.com/Eznopot/RM_api/src/Logger"
	docx "github.com/lukasjarosch/go-docx"
)

func WaitUntil(ctx context.Context, until time.Time) {
	timer := time.NewTimer(time.Until(until))
	defer timer.Stop()

	select {
	case <-timer.C:
		return
	case <-ctx.Done():
		return
	}
}

func JsonArrayToMap(jsonString string) ([]map[string]interface{}, error) {
	var jsonMap []map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &jsonMap)
	if err != nil {
		return nil, err
	}
	return jsonMap, nil
}

func CreateWordCV(id int, initial, formation, experience, competence string) string {
	replaceMap := docx.PlaceholderMap{
		"_initial_":    initial,
		"_competence_": competence,
		"_formation_":  formation,
		"_experience_": experience,
	}

	doc, err := docx.Open("./CV/template/templateCV.docx")
	if err != nil {
		return ""
	}

	err = doc.ReplaceAll(replaceMap)
	if err != nil {
		return ""
	}
	var path = "./CV/" + strconv.Itoa(id) + initial + ".docx"
	err = doc.WriteToFile(path)
	if err != nil {
		return ""
	}
	return path
}

const (
	Red    string = "\033[31m"
	Blue   string = "\033[34m"
	Yellow string = "\033[33m"
	Green  string = "\033[32m"
)

func BetterPrint(color string, elem string, noReset ...bool) {
	if noReset != nil && noReset[0] {
		println(color + elem)
	} else {
		println(color + elem + "\033[0m")
	}
}

func ResetColor() {
	println("\033[0m")
}

func IsDayOff(holliday map[string]interface{}, date time.Time) bool {
	for dayOff := range holliday {
		dayOffTime, err := time.Parse("2006-01-02", dayOff)
		if err != nil {
			logger.Error(err.Error())
			break
		}
		y1, m1, d1 := dayOffTime.Date()
		y2, m2, d2 := date.Date()
		if y1 == y2 && m1 == m2 && d1 == d2 {
			return true
		}
	}
	return false
}

func GetHollidays(year int) map[string]interface{} {
	var url string
	if year != -1 {
		url = "https://calendrier.api.gouv.fr/jours-feries/metropole/" + strconv.Itoa(year) + ".json"
	} else {
		url = "https://calendrier.api.gouv.fr/jours-feries/metropole.json"
	}

	// Create an HTTP client
	client := &http.Client{}

	// Create an HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Error(err.Error())
		return nil
	}

	// Send the request and retrieve the response
	resp, err := client.Do(req)
	if err != nil {
		logger.Error(err.Error())
		return nil
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err.Error())
		return nil
	}

	// Unmarshal the JSON response
	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		logger.Error(err.Error())
		return nil
	}
	return response
}
