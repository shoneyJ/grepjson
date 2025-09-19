package search

import "fmt"

func searchRecursive(data interface{}, keyPattern string, currentPath string, results *[]MatchResult, maxDistance int) {
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
			searchRecursive(val, keyPattern, newPath, results, maxDistance)
		}
	case []interface{}:
		for i, val := range v {
			searchRecursive(val, keyPattern, buildPath(currentPath, fmt.Sprintf("[%d]", i)), results, maxDistance)
		}

	}
}
