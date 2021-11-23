package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/sfomuseum/go-accession-numbers"
	"log"
	"os"
)

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

		var def *accessionnumbers.Definition

		dec := json.NewDecoder(fh)
		err = dec.Decode(&def)

		if err != nil {
			log.Fatalf("Failed to decode %s, %v", path, err)
		}

		err = runTests(def)

		if err != nil {
			log.Fatalf("Failed to run tests for %s, %v", def.OrganizationName, err)
		}

		log.Printf("All tests pass for %s\n", def.OrganizationName)
	}
}

func runTests(def *accessionnumbers.Definition) error {

	for _, p := range def.Patterns {

		for str, expected_results := range p.Tests {

			// no results means 'skip', for example if a particular test is being
			// problamatic

			if len(expected_results) == 0 {
				log.Printf("[%s] SKIP %s\n", def.OrganizationURI, str)
				continue
			}

			matches, err := accessionnumbers.FindMatches(str, p.Pattern)

			if err != nil {
				return fmt.Errorf("Failed to find matches for '%s' using '%s' (%s), unexpected error: %w", str, p.Pattern, def.OrganizationName, err)
			}

			expected_count := len(expected_results)
			count := len(matches)

			if count != expected_count {
				return fmt.Errorf("Failed to find matches for '%s' using '%s' (%s), expected %d matches but got %d (%v)", str, p.Pattern, def.OrganizationName, expected_count, count, matches)
			}

			for i, expected_value := range expected_results {

				if matches[i] != expected_value {
					return fmt.Errorf("Match %d failed for '%s' using '%s' (%s), expected '%s' but got '%s'", i, str, p.Pattern, def.OrganizationName, expected_value, matches[i])
				}
			}

			log.Printf("[%s] OK %s\n", def.OrganizationURI, str)
		}
	}

	return nil
}
