package utils

import (
	"encoding/json"
	"log"
)

func StringToJson() {}

func StructToJsonString(struc map[string]string) string {
	jsonResponse, err := json.Marshal(struc)
	if err != nil {
		log.Fatal(err)
	}
	return string(jsonResponse)
}
