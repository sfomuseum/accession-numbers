package accessionnumbers

import (
	"encoding/json"
	"fmt"
	"os"
)

func loadTestDefinition() (*Definition, error) {

	path := "fixtures/sfomuseum.json"

	r, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("Failed to open %s for reading, %w", path, err)
	}

	defer r.Close()

	var def *Definition

	dec := json.NewDecoder(r)
	err = dec.Decode(&def)

	if err != nil {
		return nil, fmt.Errorf("Failed to decode defintion for %s, %w", path, err)
	}

	return def, nil
}
