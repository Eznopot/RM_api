package utils

import (
	"encoding/json"
	"strconv"

	docx "github.com/lukasjarosch/go-docx"
)

func JsonArrayToMap(jsonString string) ([]map[string]interface{}, error) {
	var jsonMap []map[string]interface{}
	err := json.Unmarshal([]byte(jsonString ), &jsonMap)
	if (err != nil) {
		return nil, err;
	}
	return jsonMap, nil;
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
