package utils

import (
	"strconv"

	docx "github.com/lukasjarosch/go-docx"
)

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
