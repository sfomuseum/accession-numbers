package accessionnumbers

// type Defintion provides a struct containing accession number patterns and URIs for an organization.
type Definition struct {
	// The name of the organization associated with this definition.
	OrganizationName string `json:"organization_name"`
	// The URL of the organization associated with this definition.
	OrganizationURL string `json:"organization_url"`
	// A valid URI template (RFC 6570) used to generate the URL for an object given its accession number.
	ObjectURL string `json:"object_url,omitempty"`
	// A valid URI template (RFC 6570) used to generate the IIIF manifest URL for an object given its accession number.
	IIIFManifest string `json:"iiif_manifest,omitempty"`
	// A valid URI template (RFC 6570) used to generate an OEmbed profile URL for an object given its accession number.
	OEmbedProfile string `json:"oembed_profile,omitempty"`
	// The set of patterns used to identify and extract accession numbers associated with an organization.
	Patterns []*Pattern `json:"patterns"`
}

// type Pattern provides a struct containing patterns and tests for one or more accession numbers.
type Pattern struct {
	// The name or label for a given pattern.
	Label string `json:"label"`
	// A valid regular expression string.
	Pattern string `json:"pattern"`
	// A dictionary containing zero or more tests for validating `Pattern`. Keys contain text to extract accession numbers from and values are the list of accession numbers expected to be found in the text, in the order that they are found.
	Tests map[string][]string `json:"tests"`
}

// type Match provides a struct containing accession number details found in a body of text.
type Match struct {
	// The accession number found in a body of text.
	AccessionNumber string `json:"accession_number"`
	// The URL of the organization that the accession number is associated with. Wherever possible this should match the `OrganizationURL` property in a `Defintion` struct.
	OrganizationURL string `json:"organization,omitempty"`
}
