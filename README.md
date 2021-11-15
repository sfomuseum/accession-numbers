# accession-numbers

Machine-readable regular expressions for identifying accession numbers for cultural heritage organizations in text.

## Important

This is work in progress. Things may still change.

## Motivation

The goal of this package is to have a collection of machine-readable regular expression patterns that can be used by applications to isolate accession numbers in arbitrary bodies of text. For example these data might be used by the [sfomuseum/ios-label-whisperer](https://github.com/sfomuseum/ios-label-whisperer) application.

## Data

Data for individual organizations are defined in [`data/{organization}.json`](data) files. These files lack a well-defined schema at this time.

The simplest version of a data file consists of `name` and `url` properties identifying an organization and a `patterns` properties which contains one or more dictionaries containing regular expression patterns that can be used to isolate accession numbers in a body of text. For example:

```
{
    "name": "National Museum of African American History and Culture",
    "url": "https://nmaahc.si.edu/",
    "patterns": [
	{
	    "name": "common",
	    "pattern": "((?:\\d{4})\\.(?:\\d+)(?:\\.\\d+){0,2})",
	    "tests": {
		"2013.68.19": 1,
		"2012.110": 1,
		"2016.5.2.11": 1,
		"2014.270.2": 1
	    }
	}
    ]
}
```

### Patterns

Regular expression patterns should match the entire accession number and any interior matches should be non-greedy.

### Tests

Tests for any given pattern are defined as a dictionary whose values are strings to match against (the current pattern) and whose values are the number of expected matches for a corresponding string (key).

_Tests are run using the [cmd/test-runner](cmd/test-runner]) tool which is written in Go and uses the [regexp.FindStringSubmatch](https://pkg.go.dev/regexp#Regexp.FindStringSubmatch) method to find matches._

## Tests

This packages comes with a command-line tool for running tests against some or all the files in the `data` directory. The tool is called `test-runner` and its source code can be found in the [cmd/test-runner](cmd/test-runner) folder. It has also been pre-compiled to run on Windows, Linux and Mac OS computers. These binary versions are kept in the `bin/(YOUR-OS-HERE)` folders. Possible values for `(YOUR-OS-HERE)` are:

| Operating System | Value |
| --- | --- |
| Linux | linux |
| Mac OS | darwin |
| Windows | windows |

To run the `test-runner` tool type the following from a terminal window:

```
$> bin/(YOUR-OS-HERE)/test-runner data/*.json
```

And you should see something like this:

```
$> ./bin/darwin/test-runner data/*.json
2021/11/14 15:38:21 All tests pass for Cooper Hewitt Smithsonian National Design Museum
2021/11/14 15:38:21 All tests pass for SFO Museum
```

If you have the `make` application installed on your computer you can also simply run the `tests` Makefile target. For example:

```
$> make tests
bin/darwin/test-runner data/*.json
2021/11/14 16:19:59 All tests pass for Art Institute of Chicago
2021/11/14 16:19:59 All tests pass for Cooper Hewitt Smithsonian National Design Museum
2021/11/14 16:19:59 All tests pass for Denver Museum of Nature & Science
2021/11/14 16:19:59 All tests pass for Getty Center
2021/11/14 16:19:59 All tests pass for Metropolitan Museum of Art
2021/11/14 16:19:59 All tests pass for Museum of Modern Art
2021/11/14 16:19:59 All tests pass for National Air and Space Museum
2021/11/14 16:19:59 All tests pass for National Gallery of Art
2021/11/14 16:19:59 All tests pass for National Museum of Anthropology
2021/11/14 16:19:59 All tests pass for National Museum of African American History and Culture
2021/11/14 16:19:59 All tests pass for National Museum of American History
2021/11/14 16:19:59 All tests pass for Smithsonian National Museum of Natural History
2021/11/14 16:19:59 All tests pass for SFO Museum
```

## Help wanted

Contributions for missing organizations and corrections for existing patterns are welcome.

## Contributors

* [Bruce Wyman](http://www.usd-mach.com/)

## See also

* https://github.com/sfomuseum/ios-label-whisperer
