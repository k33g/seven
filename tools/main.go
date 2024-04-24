package tools

import (
	"bufio"
	"encoding/json"
	"os"
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

func EnvSubst(source string) (string, error) {
	env := os.Environ()
	target := ""
	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan() {
		row := scanner.Text()
		// substitute environment variables
		for _, value := range env {
			pair := strings.SplitN(value, "=", 2)
			row = strings.Replace(row, "${"+pair[0]+"}", pair[1], -1)
			row = strings.Replace(row, "$"+pair[0], pair[1], -1)
		}
		target += row + "\n"

	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return target, nil

}
