package main

import (
	"reflect"
	"testing"

	"github.com/Eznopot/RM_api/src/utils"
)

func TestJsonArrayToMap(t *testing.T) {
	type args struct {
		jsonString string
	}
	tests := []struct {
		name    string
		args    args
		want    []map[string]interface{}
		wantErr bool
	}{
		{
			name: "Test JsonArrayToMap multiple element",
			args: args{
				jsonString: "[{\"id\":1,\"name\":\"test\"},{\"id\":2,\"name\":\"test2\"}]",
			},
			want: []map[string]interface{}{
				{
					"id":   float64(1),
					"name": "test",
				},
				{
					"id":   float64(2),
					"name": "test2",
				},
			},
			wantErr: false,
		},
		{
			name: "TestJsonArray one element",
			args: args{
				jsonString: "[{\"id\":1,\"name\":\"test\"}]",
			},
			want: []map[string]interface{}{
				{
					"id":   float64(1),
					"name": "test",
				},
			},
			wantErr: false,
		},
		{
			name: "TestJsonArray empty",
			args: args{
				jsonString: "[]",
			},
			want:    []map[string]interface{}{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := utils.JsonArrayToMap(tt.args.jsonString)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonArrayToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonArrayToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
