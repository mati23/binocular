package utils

func TrucateString(stringToBeTruncated string, size int) string {
	if len(stringToBeTruncated) > 30 {
		return stringToBeTruncated[:30] + "..."
	}
	return stringToBeTruncated
}
