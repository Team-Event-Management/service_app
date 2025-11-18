package utils

import "strconv"

// ParseInt safely converts string to int with default fallback value.
// Example: ParseInt("10", 0) => 10
//          ParseInt("", 5) => 5
func ParseInt(s string, defaultVal int) int {
	if s == "" {
		return defaultVal
	}
	val, err := strconv.Atoi(s)
	if err != nil {
		return defaultVal
	}
	return val
}
