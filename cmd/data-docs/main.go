package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"github.com/sfomuseum/go-accession-numbers"
	"log"
	"os"
	_ "path/filepath"
	"sort"
	"text/template"
	"time"
)

//go:embed README.txt
var readme_t []byte

func main() {

	flag.Parse()

	t := template.New("readme")
	t, err := t.Parse(string(readme_t))

	if err != nil {
		log.Fatalf("Failed to parse README template, %v", err)
	}

	paths_lookup := make(map[string][]string)
	org_lookup := make(map[string]*accessionnumbers.Definition)

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

		paths, ok := paths_lookup[def.OrganizationName]

		if !ok {
			paths = make([]string, 0)
		}

		paths = append(paths, path)
		paths_lookup[def.OrganizationName] = paths

		org_lookup[path] = def
	}

	names := make([]string, 0)

	for n, _ := range paths_lookup {
		names = append(names, n)
	}

	sort.Strings(names)

	orgs := make([]*accessionnumbers.Definition, 0)

	for _, n := range names {

		for _, p := range paths_lookup[n] {
			o := org_lookup[p]
			// o.Path = filepath.Base(p)
			orgs = append(orgs, o)
		}
	}

	now := time.Now()

	vars := struct {
		Orgs         []*accessionnumbers.Definition
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
