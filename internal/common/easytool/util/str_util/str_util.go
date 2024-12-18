package str_util

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
