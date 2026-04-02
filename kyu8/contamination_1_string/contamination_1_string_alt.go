package contamination_1_string

import (
	"strings"
	"unicode/utf8"
)

// Kata: https://www.codewars.com/kata/596fba44963025c878000039
//
// An AI has infected a text with a character!!
// This text is now fully mutated to this character.
//
// If the text or the character are empty, return an empty string.
// There will never be a case where both are empty as nothing
// is going on!!

func ContaminationAlt(text string, char string) string {
	var b strings.Builder

	for i := 0; i < utf8.RuneCountInString(text); i++ {
		b.WriteString(char)
	}

	return b.String()
}
