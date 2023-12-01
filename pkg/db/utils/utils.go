package utils

func EmptyOr(str string, defaultStr string) string {
	if str == "" {
		return defaultStr
	}

	return str
}
