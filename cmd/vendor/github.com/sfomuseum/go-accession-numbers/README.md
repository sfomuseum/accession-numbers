# go-accession-numbers

Go package providing methods for identifying and extracting accession numbers from arbitrary bodies of text.

## Documentation

[![Go Reference](https://pkg.go.dev/badge/github.com/sfomuseum/go-accession-numbers.svg)](https://pkg.go.dev/github.com/sfomuseum/go-accession-numbers)

## Example

_Error handling omitted for the sake of brevity._

### Basic

```
package main

import (
	"fmt"
	"github.com/sfomuseum/go-accession-numbers"
)

func main() {

	re_pat := `((?:L|R)?(?:\d+)\.(?:\d+)\.(?:\d+)(?:\.(?:\d+))?(?:(?:\s?[sa-z])+)?)`
     
	texts := []string{
     		"2000.058.1185 a c",
		"This is an object\nGift of Important Donor\n1994.18.175\n\nThis is another object\nAnonymouts Gift\n1994.18.165 1994.18.199a\n2000.058.1185 a c\nOil on canvas",
     	}

	for _, t := range texts {
		     
     		m, _ := accessionnumbers.FindMatches(text, re)

		for _, m := range matches {
			fmt.Printf("%s\n", m)
		}
     	}
```

This would yield:

```
2000.058.1185 a c
1994.18.175
1994.18.165
1994.18.199a
2000.058.1185 a c
```

### Using "defintion" files

"Definition" files are provided by the [sfomuseum/accession-numbers](https://github.com/sfomuseum/accession-numbers) package.

_Note: As of this writing this package assumes that definition file structure that hasn't been adopted by the `main` branch of the `accession-numbers` package yet._

```
package main

import (
	"encoding/json"
	"fmt"		
	"github.com/sfomuseum/go-accession-numbers"
	"os"
)

func main() {

	var def *Definition
	
	r, _ := os.Open("fixtures/sfomuseum.json")

	dec := json.NewDecoder(r)
	dec.Decode(&def)

	re_pat := `((?:L|R)?(?:\d+)\.(?:\d+)\.(?:\d+)(?:\.(?:\d+))?(?:(?:\s?[sa-z])+)?)`
     
	texts := []string{
     		"2000.058.1185 a c",
		"This is an object\nGift of Important Donor\n1994.18.175\n\nThis is another object\nAnonymouts Gift\n1994.18.165 1994.18.199a\n2000.058.1185 a c\nOil on canvas",
     	}

	for _, t := range texts {

		matches, _ := accessionnumbers.ExtractFromText(t, def)
		
		for _, m := range matchess {
			fmt.Printf("%s (%s)\n", m.AccessionNumber, m.OrganizationURL)
		}
     	}
```

This would yield:

```
2000.058.1185 a c (https://sfomuseum.org/)
1994.18.175 (https://sfomuseum.org/)
1994.18.165 (https://sfomuseum.org/)
1994.18.199a (https://sfomuseum.org/)
2000.058.1185 a c (https://sfomuseum.org/)
```

## See also

* https://github.com/sfomuseum/accession-numbers