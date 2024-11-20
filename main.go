package numutils

import (
	"fmt"
	"strconv"
	"strings"
)

func IsNumber(s string) bool {
	// Trim spaces from the string
	s = strings.TrimSpace(s)

	// List of non-numeric special cases
	nonNumericCases := map[string]bool{
		"Infinity":  true,
		"-Infinity": true,
		"NaN":       true,
	}

	// Check for non-numeric special cases
	if nonNumericCases[s] {
		return false
	}

	// Check if the string is a hexadecimal number
	if strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X") {
		if _, err := strconv.ParseInt(s[2:], 16, 64); err == nil {
			return true
		}
	}

	// Try to parse the string as a float
	if _, err := strconv.ParseFloat(s, 64); err == nil {
		return true
	}

	// Try to parse the string as an integer
	if _, err := strconv.ParseInt(s, 10, 64); err == nil {
		return true
	}

	// If all parsing attempts fail, return false
	return false
}

func main() {
	fmt.Println(IsNumber("123"))
}
