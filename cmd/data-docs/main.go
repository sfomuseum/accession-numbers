package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"log"
	"os"
	"path/filepath"
	"sort"
	"text/template"
	"time"
)

//go:embed README.txt
var readme_t []byte

// START OF please reconcile me with cmd/test-runner

type Organization struct {
	Name     string     `json:"name`
	URL      string     `json:"url"`
	Patterns []*Pattern `json:"patterns"`
	Path     string     `json:"path,omitempty"`
}

type Pattern struct {
	Name    string         `json:"name"`
	Pattern string         `json:"pattern"`
	Tests   map[string]int `json:"tests"`
}

// END OF please reconcile me with cmd/test-runner

func main() {

	flag.Parse()

	t := template.New("readme")
	t, err := t.Parse(string(readme_t))

	if err != nil {
		log.Fatalf("Failed to parse README template, %v", err)
	}

	paths_lookup := make(map[string][]string)
	org_lookup := make(map[string]*Organization)

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

		paths, ok := paths_lookup[org.Name]

		if !ok {
			paths = make([]string, 0)
		}

		paths = append(paths, path)
		paths_lookup[org.Name] = paths

		org_lookup[path] = org
	}

	names := make([]string, 0)

	for n, _ := range paths_lookup {
		names = append(names, n)
	}

	sort.Strings(names)

	orgs := make([]*Organization, 0)

	for _, n := range names {

		for _, p := range paths_lookup[n] {
			o := org_lookup[p]
			o.Path = filepath.Base(p)
			orgs = append(orgs, o)
		}
	}

	now := time.Now()

	vars := struct {
		Orgs         []*Organization
		LastModified time.Time
	}{
		Orgs:         orgs,
		LastModified: now,
	}

	err = t.Execute(os.Stdout, vars)

	if err != nil {
		log.Fatalf("Failed to render template, %v", err)
	}
}
