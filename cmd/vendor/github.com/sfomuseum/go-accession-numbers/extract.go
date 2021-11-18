package accessionnumbers

import (
	"fmt"
	"regexp"
	"strings"
)

// Extract a list of accession numbers (as `Match` instances) from text, using regular expressions defined in defs.
func ExtractFromText(text string, defs ...*Definition) ([]*Match, error) {

	matches := make([]*Match, 0)

	for _, d := range defs {

		md, err := ExtractFromTextWithDefinition(text, d)

		if err != nil {
			return nil, err
		}

		for _, m := range md {
			matches = append(matches, m)
		}
	}

	return matches, nil
}

// Extract a list of accession numbers (as `Match` instances) from text, using regular expressions defined in def.
func ExtractFromTextWithDefinition(text string, def *Definition) ([]*Match, error) {

	matches := make([]*Match, 0)

	for _, p := range def.Patterns {

		mp, err := ExtractFromTextWithPattern(text, p)

		if err != nil {
			return nil, err
		}

		for _, m := range mp {
			matches = append(matches, m)
		}

	}

	return matches, nil
}

// Extract a list of accession numbers (as `Match` instances) from text, using regular expressions defined in pat.
func ExtractFromTextWithPattern(text string, pat *Pattern) ([]*Match, error) {

	matches := make([]*Match, 0)

	mp, err := FindMatches(text, pat.Pattern)

	if err != nil {
		return nil, err
	}

	for _, str := range mp {
		m := &Match{AccessionNumber: str}
		matches = append(matches, m)
	}

	return matches, nil
}

// Extract a list of accession numbers (as strings) from text, using a regular expression pattern defined by pat.
func FindMatches(text string, pat string) ([]string, error) {

	// Specifically we are looking for accession numbers at the end
	// of a buffer

	re_pat := fmt.Sprintf(".*?%s", pat)
	re, err := regexp.Compile(re_pat)

	if err != nil {
		return nil, fmt.Errorf("Failed to compile pattern (%s), %w", re_pat, err)
	}

	matches := make([]string, 0)

	// Just get rid of newlines to start with because they get parsed in to separate
	// '\' and 'n' runes by Go.

	text = strings.Replace(text, `\n`, " ", -1)
	buf := ""

	seen := ""

	for _, rune := range text {

		char := string(rune)
		seen += char

		switch char {
		case " ":

			found := find(buf, re)

			// No matches so keep adding to buf - this might happen
			// with accession numbers like 'Obj: 96681' which really
			// does exist (AIC)

			if len(found) == 0 {
				buf += char
				continue
			}

			// In order to account for things like `2000.058.1185 a c` (sfomuseum)
			// we need to continue read ahead testing buf until it *doesn't* match.
			// That is, given `2000.058.1185 a c`:
			// `2000.058.1185`       matches
			// `2000.058.1185 `      matches
			// `2000.058.1185 a`     matches
			// `2000.058.1185 a`     matches
			// `2000.058.1185 a `    matches
			// `2000.058.1185 a c`   matches
			// `2000.058.1185 a c `  matches
			// `2000.058.1185 a c (` does not match

			remaining := strings.Replace(text, seen, "", 1)

			buf += char

			for _, r := range remaining {

				buf += string(r)
				found_more := find(buf, re)

				if len(found_more) == 0 {
					break
				}

				found = found_more
			}

			for _, m := range found {
				matches = append(matches, m)
			}

			buf = ""

		default:
			buf = fmt.Sprintf("%s%s", buf, char)
		}

	}

	if buf != "" {

		for _, m := range find(buf, re) {
			matches = append(matches, m)
		}
	}

	// log.Printf("MATCHES '%s', %v\n", text, matches)
	return matches, nil
}

func find(buf string, re *regexp.Regexp) []string {

	m := re.FindStringSubmatch(buf)

	if len(m) <= 1 {
		return []string{}
	}

	return m[1:]
}
