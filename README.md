# accession-numbers

Machine-readable regular expressions for identifying accession numbers for cultural heritage organizations in text.

## Important

This is work in progress. Things may still change.

## Motivation

The goal of this package is to have a collection of machine-readable regular expression patterns that can be used by applications to isolate accession numbers in arbitrary bodies of text. For example these data might be used by the [sfomuseum/ios-label-whisperer](https://github.com/sfomuseum/ios-label-whisperer) application.

## Data

Data for individual organizations are defined in `data/{organization}.json` files. These files lack a well-defined schema at this time.

The simplest version of a data file consists of `name` and `url` properties identifying an organization and a `patterns` properties which contains one or more dictionaries containing regular expression patterns that can be used to isolate accession numbers in a body of text. For example:

```
{
    "name": "SFO Museum",
    "url": "https://sfomuseum.org/",
    "patterns": [
	{
	    "name": "common",
	    "pattern": "((?:\\d+)\\.(?:\\d+)\\.(?:\\d+))",
	    "tests": {
		"1994.18.175": 1
	    }
	}
    ]
}
```

## Tests

This packages comes with a command-line tool for running tests against some or all the files in the `data` directory. The tool is called `test-runner` and its source code can be found in the [cmd/test-runner](cmd/test-runner) folder. It has also been pre-compiled to run on Windows, Linux and Mac OS computers. (These binary versions are kept in the `bin/(OS)` folder.)

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

Possible values for `(YOUR-OS-HERE)` are:

| Operating System | Value |
| --- | --- |
| Linux | linux |
| Mac OS | darwin |
| Windows | windows |

If you have the `make` application installed on your computer you can also simply run the `tests` Makefile target. For example:

```
$> make tests
bin/darwin/test-runner data/*.json
2021/11/14 15:40:20 All tests pass for Cooper Hewitt Smithsonian National Design Museum
2021/11/14 15:40:20 All tests pass for SFO Museum
```


## Help wanted

Contributions for missing organizations and corrections for existing patterns are welcome.

## See also

* https://github.com/sfomuseum/ios-label-whisperer