package utils

import "encoding/json"

// ValidateJSON checks if the given string is a valid JSON
func ValidateJSON(input string) error {
	var js map[string]interface{}
	return json.Unmarshal([]byte(input), &js)
}
