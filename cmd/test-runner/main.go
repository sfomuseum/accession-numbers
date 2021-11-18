package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// START OF please reconcile me with cmd/data-docs

type Organization struct {
	Name     string     `json:"name`
	URL      string     `json:"url"`
	IIIFManifest string `json:"iiif_manifest"`
	OEmbedProfile string `json:"oembed_profile"`
	Patterns []*Pattern `json:"patterns"`
}

type Pattern struct {
	Name    string         `json:"name"`
	Pattern string         `json:"pattern"`
	Tests   map[string][]string `json:"tests"`
}

// END OF please reconcile me with cmd/data-docs

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Parse one or more data JSON files with accession number patterns ensuring valid regular expressions and running all defined tests.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s org.json(N) org.json(N)\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	for _, path := range flag.Args() {

		fh, err := os.Open(path)

		if err != nil {
			log.Fatalf("Failed to open %s, %v", path, err)
		}

		defer fh.Close()

		var org *Organization

		dec := json.NewDecoder(fh)
		err = dec.Decode(&org)

		if err != nil {
			log.Fatalf("Failed to decode %s, %v", path, err)
		}

		err = runTests(org)

		if err != nil {
			log.Fatalf("Failed to run tests for %s, %v", org.Name, err)
		}

		log.Printf("All tests pass for %s\n", org.Name)
	}
}

func runTests(org *Organization) error {

	for _, p := range org.Patterns {

		for str, expected_results := range p.Tests {

			// no results means 'skip', for example if a particular test is being
			// problamatic
			
			if len(expected_results) == 0 {
				log.Printf("[%s] SKIP %s\n", org.URL, str)				
				continue
			}

			matches, err := findMatches(str, p.Pattern)

			if err != nil {
				return fmt.Errorf("Failed to find matches for '%s' using '%s' (%s), unexpected error: %w", str, p.Pattern, org.Name, err)
			}

			expected_count := len(expected_results)
			count := len(matches)

			if count != expected_count {
				return fmt.Errorf("Failed to find matches for '%s' using '%s' (%s), expected %d matches but got %d (%v)", str, p.Pattern, org.Name, expected_count, count, matches)
			}

			for i, expected_value := range expected_results {

				if matches[i] != expected_value {
					return fmt.Errorf("Match %d failed for '%s' using '%s' (%s), expected '%s' but got '%s'", i, str, p.Pattern, org.Name, expected_value, matches[i])
				}
			}

			log.Printf("[%s] OK %s\n", org.URL, str)
		}
	}

	return nil
}

// START OF put me in a go-accession-numbers package

func findMatches(text string, pat string) ([]string, error) {

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

// END OF put me in a go-accession-numbers package
