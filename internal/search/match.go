package search

import "strings"

// MatchResult represents a found match with path and value
type MatchResult struct {
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
}

func buildPath(base, part string) string {
	if strings.HasPrefix(part, "[") {
		return base + part
	}
	if base == "" {
		return part
	}
	return base + "." + part
}
