#!/bin/bash
set -e

# Installation parameters
VERSION="1.0"
APP_NAME="grepjson"
BINARY_URL="https://github.com/shoneyJ/grepjson/releases/download/v${VERSION}/${APP_NAME}"
INSTALL_DIR="/usr/local/bin"
TEMP_FILE=$(mktemp)

# Detect system architecture
ARCH=$(uname -m)
case $ARCH in
x86_64) ARCH="amd64" ;;
arm64 | aarch64) ARCH="arm64" ;;
*)
  echo "Unsupported architecture: $ARCH"
  exit 1
  ;;
esac

# Detect OS and set binary name
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
BINARY_NAME="${APP_NAME}-${OS}-${ARCH}"

echo "Downloading ${APP_NAME} v${VERSION}..."
if ! curl -sSL -o "$TEMP_FILE" "https://github.com/shoneyJ/grepjson/releases/download/v${VERSION}/${BINARY_NAME}"; then
  echo "Error: Download failed"
  exit 1
fi

echo "Verifying and installing..."
chmod +x "$TEMP_FILE"

# Install to system directory (requires sudo)
if [ -w "$INSTALL_DIR" ]; then
  mv "$TEMP_FILE" "${INSTALL_DIR}/${APP_NAME}"
else
  echo "Requires sudo to install to ${INSTALL_DIR}"
  sudo mv "$TEMP_FILE" "${INSTALL_DIR}/${APP_NAME}"
  sudo chown root:root "${INSTALL_DIR}/${APP_NAME}"
fi

# Verify installation
if command -v "$APP_NAME" >/dev/null 2>&1; then
  echo "Successfully installed:"
  "$APP_NAME" --version || echo "Version check not available"
else
  echo "Installation failed - please ensure ${INSTALL_DIR} is in your PATH"
  exit 1
fi
