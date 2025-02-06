package str_util

import "strings"

// IsBlank check if string is blank
func IsBlank(str string) bool {
	if str == "" {
		return true
	}
	return false
}

// IsNotBlank check if string is not blank
func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

// ExtractStringFromBorder extracts the first substring found between the border
func ExtractStringFromBorder(input, left, right string) string {
	leftIndex := strings.Index(input, left)
	rightIndex := strings.Index(input, right)
	if leftIndex == -1 || rightIndex == -1 {
		return ""
	}
	startIndex := leftIndex + len(left)
	if rightIndex < startIndex {
		return ""
	}
	return input[startIndex:rightIndex]
}
