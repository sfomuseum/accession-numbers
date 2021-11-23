package accessionnumbers

import (
	"fmt"
	"github.com/jtacoma/uritemplates"
)

// type Definition provides a struct containing accession number patterns and URIs for an organization.
type Definition struct {
	// The name of the organization associated with this definition.
	OrganizationName string `json:"organization_name"`
	// The URI of the organization associated with this definition.
	OrganizationURI string `json:"organization_url"`
	// A valid URI template (RFC 6570) used to generate the URI for an object given its accession number.
	ObjectURITemplate string `json:"object_url,omitempty"`
	// A valid URI template (RFC 6570) used to generate the IIIF manifest URI for an object given its accession number.
	IIIFManifestTemplate string `json:"iiif_manifest,omitempty"`
	// A valid URI template (RFC 6570) used to generate an OEmbed profile URI for an object given its accession number.
	OEmbedProfileTemplate string `json:"oembed_profile,omitempty"`
	// A valid Who's On First ID representing the organization.
	WhosOnFirstId int64 `json:"whosonfirst_id,omitempty"`
	// The set of patterns used to identify and extract accession numbers associated with an organization.
	Patterns []*Pattern `json:"patterns"`
}

// IIIFManifestURI returns a IIIF manifest URI for accession_number, assuming a corresponding URI template exists in def.
func (def *Definition) IIIFManifestURI(accession_number string) (string, error) {

	if def.IIIFManifestTemplate == "" {
		return "", fmt.Errorf("IIIFManifestURITemplate is undefined")
	}

	return def.expandURITemplate(def.IIIFManifestTemplate, accession_number)
}

// IIIFManifestURI returns an OEmbed profile URI for accession_number, assuming a corresponding URI template exists in def.
func (def *Definition) OEmbedProfileURI(accession_number string) (string, error) {

	if def.OEmbedProfileTemplate == "" {
		return "", fmt.Errorf("OEmbedProfileURITemplate is undefined")
	}

	return def.expandURITemplate(def.OEmbedProfileTemplate, accession_number)
}

// IIIFManifestURI returns a object URI for accession_number, assuming a corresponding URI template exists in def.
func (def *Definition) ObjectURI(accession_number string) (string, error) {

	if def.ObjectURITemplate == "" {
		return "", fmt.Errorf("ObjectURITemplate is undefined")
	}

	return def.expandURITemplate(def.ObjectURITemplate, accession_number)
}

func (def *Definition) expandURITemplate(str_t string, accession_number string) (string, error) {

	t, err := uritemplates.Parse(str_t)

	if err != nil {
		return "", fmt.Errorf("Failed to parse URI template, %w", err)
	}

	values := map[string]interface{}{
		"accession_number": accession_number,
	}

	str_uri, err := t.Expand(values)

	if err != nil {
		return "", fmt.Errorf("Failed to expand URI templates, %w", err)
	}

	return str_uri, nil
}
