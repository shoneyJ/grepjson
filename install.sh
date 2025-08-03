#!/usr/bin/env bash

set -e

# Variables
URL="https://github.com/shoneyJ/grepjson/releases/download/v1.0/grepjson-linux-x86_64.tar.gz"
TMP_DIR="$(mktemp -d)"
ARCHIVE_NAME="grepjson-linux-x86_64.tar.gz"
INSTALL_DIR="/usr/local/bin"
BINARY_NAME="grepjson"

echo "📦 Downloading grepjson..."
curl -L "$URL" -o "$TMP_DIR/$ARCHIVE_NAME"

echo "📂 Extracting..."
tar -xzf "$TMP_DIR/$ARCHIVE_NAME" -C "$TMP_DIR"

echo "🚚 Installing to $INSTALL_DIR..."
sudo mv "$TMP_DIR/$BINARY_NAME" "$INSTALL_DIR/$BINARY_NAME"
sudo chmod +x "$INSTALL_DIR/$BINARY_NAME"

echo "✅ Installed grepjson at $INSTALL_DIR/$BINARY_NAME"
echo "🔍 Version:"
"$BINARY_NAME" --help || echo "Run '$BINARY_NAME' to test it."

# Clean up
rm -rf "$TMP_DIR"
