package format

import (
	"strings"
	"unicode"
)

func AddBRNewLine(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "<br>", "\n"), "<br/>", "\n")
}

// probably highly inefficient, but oh well
func ReplaceAllOverlap(s string, old string, new string) string {
	for strings.Contains(s, old) {
		s = strings.ReplaceAll(s, old, new)
	}
	return s
}

func RemoveSpecialCharacters(s string) string {
	return strings.TrimFunc(s, func(r rune) bool {
		return unicode.IsLetter(r) || unicode.IsSpace(r) || unicode.IsNumber(r)
	})
}

func FormatNewLine(s string) string {
	return strings.TrimSpace(ReplaceAllOverlap(s, "\n\n\n", "\n\n"))
}

func FormatSongString(s string) string {
	return RemoveSpecialCharacters(ReplaceAllOverlap(strings.ToLower(strings.TrimSpace(s)), "  ", " "))
}
