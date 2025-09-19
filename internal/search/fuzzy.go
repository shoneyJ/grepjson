package search

import (
	"strings"

	"github.com/texttheater/golang-levenshtein/levenshtein"
)

func fuzzyMatch(key, pattern string, maxDistance int) bool {
	if pattern == "" {
		return false
	}
	// Exact match takes precedence
	if strings.Contains(key, pattern) {
		return true
	}
	// Compute the Levenshtein distance for fuzzy matching
	dist := levenshtein.DistanceForStrings([]rune(key), []rune(pattern), levenshtein.DefaultOptions)
	return dist <= maxDistance
}
