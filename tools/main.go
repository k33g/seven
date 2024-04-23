package tools

import (
	"encoding/json"
	"strings"
)

// ✋ this is not used
func CleanJsonString(jsonString string) (string, error) {

	stringToClean := strings.ReplaceAll(jsonString, "\\n", "")
	stringToClean = strings.ReplaceAll(stringToClean, "\\\"", "\"")
	// ✋ it does not work with UTF8
	stringToClean = stringToClean[1:]
	stringToClean = stringToClean[:len(stringToClean)-1]
	cleanMap := map[string]any{}
	err := json.Unmarshal([]byte(stringToClean), &cleanMap)
	if err != nil {
		return "", err
	}
	return AnyToJsonString(cleanMap)
}

// ✋ this is not used
func AnyToJsonString(jsonAny any) (string, error) {
	bytesToClean, err := json.Marshal(jsonAny)
	if err != nil {
		return "", err
	}
	return string(bytesToClean), nil
}
