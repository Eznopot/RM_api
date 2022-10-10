package main

import (
	"reflect"
	"testing"

	"github.com/Eznopot/RM_api/src/utils"
)

func TestJsonArrayToMap(t *testing.T) {
	jsonArray := "[{\"id\":1,\"name\":\"test\"},{\"id\":2,\"name\":\"test2\"}]"
	expectedListMap := []map[string]interface{}{
		{
			"id": 1,
			"name": "test",
		},
		{
			"id": 2,
			"name": "test2",
		},
	}
	listMap, err := utils.JsonArrayToMap(jsonArray)
	if (err != nil) {
		t.Errorf("JsonArrayToMap() error = %v, wantErr %v", err, false)
		return;
	}
	
    if reflect.DeepEqual(listMap,  expectedListMap){
        t.Errorf("Expected map(%s) is not same as"+
         " actual map (%s)", expectedListMap, listMap)
    }
}