package search

import (
	"encoding/json"
	"fmt"
)

// grepJSON searches a JSON document for keys that match the given pattern
func GrepJSON(jsonBytes []byte, keyPattern string, maxDistance int) ([]MatchResult, error) {
	var data interface{}
	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	results := searchConcurrent(data, keyPattern, maxDistance)

	return results, nil
}
