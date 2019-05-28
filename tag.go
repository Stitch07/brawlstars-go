package brawlstars

import "strings"

const validChars = "0289PYLQGRJCUV"

func ValidateTag(tag string) (string, bool) {
	// remove leading #
	tag = strings.TrimPrefix(tag, "#")
	// valid tags have to be at least 3 characters long
	if len(tag) < 3 {
		return "", false
	}
	for _, char := range []rune(tag) {
		if !strings.ContainsRune(validChars, char) {
			return "", false
		}
	}
	return tag, true
}
