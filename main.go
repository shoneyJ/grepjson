package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/texttheater/golang-levenshtein/levenshtein"
)

// MatchResult represents a found match with path and value
type MatchResult struct {
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
}

// grepJSON searches a JSON document for keys that match the given pattern
func grepJSON(jsonBytes []byte, keyPattern string, maxDistance int) ([]MatchResult, error) {
	var data interface{}
	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	results := make([]MatchResult, 0)
	search(data, keyPattern, "", &results, maxDistance)

	return results, nil
}

// search recursively traverses the JSON structure
func search(data interface{}, keyPattern string, currentPath string, results *[]MatchResult, maxDistance int) {
	switch v := data.(type) {
	case map[string]interface{}:
		for k, val := range v {
			newPath := buildPath(currentPath, k)

			if fuzzyMatch(k, keyPattern, maxDistance) {
				*results = append(*results, MatchResult{
					newPath,
					val,
				})
			}
			search(val, keyPattern, newPath, results, maxDistance)
		}
	case []interface{}:
		for i, val := range v {
			search(val, keyPattern, buildPath(currentPath, fmt.Sprintf("[%d]", i)), results, maxDistance)
		}

	}
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

func createMatchResult(path, key string, value interface{}) MatchResult {
	return MatchResult{
		Path:  path,
		Value: value,
	}
}

func getType(value interface{}) string {
	switch value.(type) {
	case map[string]interface{}:
		return "object"
	case []interface{}:
		return "array"
	case string:
		return "string"
	case float64:
		return "number"
	case bool:
		return "boolean"
	case nil:
		return "null"
	default:
		return "unknown"
	}
}

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

func main() {
	keyPattern := flag.String("pattern", "", "Pattern to search for in the JSON keys")
	distance := flag.Int("distance", 1, "Maximum Levenshtein distance for fuzzy matching")
	flag.Parse()

	jsonData, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	results, err := grepJSON(jsonData, *keyPattern, *distance)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error searching JSON: %v\n", err)
		os.Exit(1)
	}

	output, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error formatting results: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}
