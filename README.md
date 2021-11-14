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

## Help wanted

Contributions for missing organizations and corrections for existing patterns are welcome.

## See also

* https://github.com/sfomuseum/ios-label-whisperer