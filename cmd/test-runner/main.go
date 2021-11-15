package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
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

		re, err := regexp.Compile(p.Pattern)

		if err != nil {
			return fmt.Errorf("Failed to compile pattern '%s', %w", p.Name, err)
		}

		for str, expected := range p.Tests {

			m := re.FindStringSubmatch(str)
			count := len(m)

			if count == 0 {
				return fmt.Errorf("String '%s' failed to match pattern '%s'", str, p.Name)
			}

			if (count - 1) != expected {
				return fmt.Errorf("String '%s' failed to match pattern '%s' with bad count, expected %d but got %d", str, p.Name, expected, (count - 1))
			}
		}
	}

	return nil
}
