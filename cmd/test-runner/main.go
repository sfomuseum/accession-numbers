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
	Patterns []*Pattern `json:"patterns"`
}

type Pattern struct {
	Name    string         `json:"name"`
	Pattern string         `json:"pattern"`
	Tests   map[string]int `json:"tests"`
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

		for str, expected_count := range p.Tests {

			// -1 means "to skip", for example if a particular test is being
			// problamatic
			
			if expected_count == -1 {
				continue
			}

			matches, err := findMatches(str, p.Pattern)

			if err != nil {
				return fmt.Errorf("Failed to find matches for '%s' using '%s' (%s), unexpected error: %w", str, p.Pattern, org.Name, err)
			}

			count := len(matches)

			if count != expected_count {
				return fmt.Errorf("Failed to find matches for '%s' using '%s' (%s), expected %d matches but got %d (%v)", str, p.Pattern, org.Name, expected_count, count, matches)
			}
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

	for _, rune := range text {

		char := string(rune)
		switch char {
		case " ":

			found := find(buf, re)

			// No matches so keep adding to buf - this might happen
			// with accession numbers like 'Obj: 96681' which really
			// does exist (AIC)
			
			if len(found) == 0 {
				buf = fmt.Sprintf("%s%s", buf, char)
				continue
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
