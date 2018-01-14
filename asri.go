package arstem

import (
	"log"
	"regexp"
	"strings"
)

var Debug *log.Logger
var rStripDiacritics *regexp.Regexp

func init() {
	rStripDiacritics = regexp.MustCompile("\\p{M}")
}

func StemASRI(word string) string {
	return stemSuffix(
		stemPrefix(
			stripDiacritics(word)))
}

func stripDiacritics(word string) string {
	return rStripDiacritics.ReplaceAllString(word, "")
}

func stemPrefix(word string) string {
	prefixes := []string{"وال", "ولل", "بال", "كال"}

	for _, p := range prefixes {
		if strings.HasPrefix(word, p) {
			return word[len(p):]
		}
	}

	return word
}

func stemSuffix(word string) string {
	suffixes := []string{"كما", "تين", "تان", "هما", "تما"}

	for _, p := range suffixes {
		if strings.HasSuffix(word, p) {
			return word[:len(p)]
		}
	}
	return word
}
