package swedish

import (
	"unicode/utf8"

	"github.com/kljensen/snowball/snowballword"
)

// Step 2: Search for one of the following suffixes in R1,
// and if found delete the last letter.
func step2(w *snowballword.SnowballWord) bool {

	suffix := w.FirstSuffix(
		"dd", "gd", "nn", "dt", "gt", "kt", "tt",
	)
	sLen := utf8.RuneCountInString(suffix)

	// If it is not in R1, do nothing
	if suffix == "" || sLen > len(w.RS)-w.R1start {
		return false
	}
	w.RemoveLastNRunes(1)
	return true
}
