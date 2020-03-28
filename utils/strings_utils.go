package utils

func DeleteEmpty(values []string) []string {
	var result []string
	for _, str := range values {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}
