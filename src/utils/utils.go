package utils

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

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

