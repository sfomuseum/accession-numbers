package accessionnumbers

// type Match provides a struct containing accession number details found in a body of text.
type Match struct {
	// The accession number found in a body of text.
	AccessionNumber string `json:"accession_number"`
	// The URL of the organization that the accession number is associated with. Wherever possible this should match the `OrganizationURL` property in a `Defintion` struct.
	OrganizationURL string `json:"organization,omitempty"`
}
