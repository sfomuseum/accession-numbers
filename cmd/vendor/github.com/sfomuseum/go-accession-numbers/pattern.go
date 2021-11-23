package accessionnumbers

// type Pattern provides a struct containing patterns and tests for one or more accession numbers.
type Pattern struct {
	// The name or label for a given pattern.
	Label string `json:"label"`
	// A valid regular expression string.
	Pattern string `json:"pattern"`
	// A dictionary containing zero or more tests for validating `Pattern`. Keys contain text to extract accession numbers from and values are the list of accession numbers expected to be found in the text, in the order that they are found.
	Tests map[string][]string `json:"tests"`
}
