package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/shoneyj/grepjson/internal/search"
)

func main() {
	keyPattern := flag.String("pattern", "", "Pattern to search for in the JSON keys")
	distance := flag.Int("distance", 1, "Maximum Levenshtein distance for fuzzy matching")
	flag.Parse()

	jsonData, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	results, err := search.GrepJSON(jsonData, *keyPattern, *distance)
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
