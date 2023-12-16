package utils

func PtrToString(str *string) string {
	if str == nil {
		return ""
	}

	return *str
}

func EmptyOr(str string, defaultStr string) string {
	if str == "" {
		return defaultStr
	}

	return str
}
