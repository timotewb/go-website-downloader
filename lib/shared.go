package lib

import (
	"fmt"
	"strings"
)

func StringInSlice(a string, list ...string) bool {
	for _, b := range list {
		if b == a {
			fmt.Println("  ", a, " : ", b)
			return true
		}
	}
	return false
}

func CreateURL(parts ...string) string {
	var combinedURL strings.Builder

	// Start with the first part as the base URL
	combinedURL.WriteString(parts[0])

	// Iterate over the remaining parts
	for _, part := range parts[1:] {
		// Trim leading and trailing slashes
		trimmedPart := strings.Trim(part, "/")

		// Append the next part, ensuring a single slash separates segments
		combinedURL.WriteString("/")
		combinedURL.WriteString(trimmedPart)
	}

	return combinedURL.String()
}
