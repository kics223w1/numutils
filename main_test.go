package numutils

import (
	"testing"
)

func TestIsNumber(t *testing.T) {
	// Test cases that should be recognized as numbers
	validNumbers := []string{
		"0xff",
		"5e3",
		"0",
		"0.1",
		"-0.1",
		"-1.1",
		"37",
		"3.14",
		"1",
		"1.1",
		"10",
		"10.1",
		"100",
		"-100",
		"0.1",
		"-0.1",
		"-1.1",
		"0",
		"012",
		"1",
		"1.1",
		"10",
		"10.10",
		"100",
		"5e3",
		"   56\r\n  ",
	}

	for _, num := range validNumbers {
		if !IsNumber(num) {
			t.Errorf("Expected %v to be a number", num)
		}
	}

	// Test cases that should not be recognized as numbers
	invalidNumbers := []string{
		"   ",
		"\r\n\t",
		"",
		"3a",
		"abc",
		"false",
		"null",
		"true",
		"undefined",
		"abc",
		"/foo/",
		"[1, 2, 4]",
		"Infinity",
		"-Infinity",
		"NaN",
		"{ a: 1 }",
		"{}",
		"/foo/",
		"[1, 2, 3]",
		"[1]",
		"[]",
		"true",
		"false",
		"function() {}",
		"Math.sin",
		"NaN",
		"new Date()",
		"null",
		"undefined",
		"{}",
	}

	for _, num := range invalidNumbers {
		if IsNumber(num) {
			t.Errorf("Expected %v to not be a number", num)
		}
	}
}
