# grepjson

A command-line tool to search for JSON keys matching a given pattern with optional fuzzy matching based on Levenshtein distance.

## Installation

To install `grepjson`, run the following command:


## 📦 Installation

### Quick Install (Linux/macOS)
Run this one-line command:

```bash
curl -sSL https://raw.githubusercontent.com/shoneyJ/grepjson/master/install.sh | bash
```

## Usage

### Basic Usage

To search for JSON keys matching a specific pattern, use the `-pattern` flag:

```sh
grepjson -pattern "example" < input.json
```

This will output the matches in JSON format to standard output.

### Fuzzy Matching

To enable fuzzy matching with a maximum Levenshtein distance of 2, use the `-distance` flag:

```sh
cat ./data.json | grepjson --pattern gen
```

### Help

For more information on available flags and usage, run:

```sh
grepjson --help
```


