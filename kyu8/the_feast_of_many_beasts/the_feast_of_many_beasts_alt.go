package the_feast_of_many_beasts

import "unicode/utf8"

// FeastAlt is an alternative solution using utf8 package.
// Handles unicode characters correctly, unlike the byte-index approach.
func FeastAlt(beast string, dish string) bool {
	bf, _ := utf8.DecodeRuneInString(beast)
	df, _ := utf8.DecodeRuneInString(dish)
	bl, _ := utf8.DecodeLastRuneInString(beast)
	dl, _ := utf8.DecodeLastRuneInString(dish)
	return bf == df && bl == dl
}
