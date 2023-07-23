package address

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// AddressType checks if address is a valid type and returns it.
func AddressType(address string) string {
	validTypes := []string{"rua", "avenida", "estrada", "rodovia"}

	sanitizedWord := strings.ToLower(address)
	firstWord := strings.Split(sanitizedWord, " ")[0]

	title := cases.Title(language.English)

	if includes(validTypes, firstWord) {
		return title.String(firstWord)
	}

	return "Invalid"
}

func includes[T comparable](slice []T, value T) bool {
	for _, i := range slice {
		if i == value {
			return true
		}
	}
	return false
}
