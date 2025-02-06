package slice_utils

import (
	"fmt"
	"strings"
)

// IsNotEmpty checks if the slice is not empty
func IsNotEmpty[T any](arr *[]T) bool {
	if arr != nil && len(*arr) > 0 {
		return true
	}
	return false
}

// IsEmpty checks if the slice is not empty
func IsEmpty[T any](arr *[]T) bool {
	return !IsNotEmpty(arr)
}

// ToStringSplitByComma converts the  columns to string
func ToStringSplitByComma[T any](arr *[]T) string {
	isEmpty := IsNotEmpty(arr)
	if isEmpty {
		return ""
	}
	builder := strings.Builder{}
	builder.WriteString("")

	for i, item := range *arr {
		builder.WriteString(fmt.Sprintf("%v", item))
		if i < len(*arr)-1 {
			builder.WriteString(",")
		}
	}

	return builder.String()
}
