#!/bin/bash

set -e

REPO="yourusername/cchat"  # Replace with your GitHub repo
VERSION="v0.1.0"           # Replace with your release tag
APP_NAME="cchat"
INSTALL_PATH="/usr/local/bin"

echo "Detecting system..."

OS=$(uname -s)
ARCH=$(uname -m)

case "$OS" in
  Linux*) GOOS="linux" ;;
  Darwin*) GOOS="darwin" ;;
  *) echo "Unsupported OS: $OS"; exit 1 ;;
esac

case "$ARCH" in
  x86_64) GOARCH="amd64" ;;
  arm64|aarch64) GOARCH="arm64" ;;
  *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

FILE_NAME="${APP_NAME}-${GOOS}-${GOARCH}.tar.gz"
URL="https://github.com/${REPO}/releases/download/${VERSION}/${FILE_NAME}"

echo "Downloading $FILE_NAME..."
curl -L --progress-bar "$URL" -o "$FILE_NAME"

echo "Extracting..."
tar -xzf "$FILE_NAME"

echo "Installing binary to $INSTALL_PATH..."
chmod +x "$APP_NAME"
sudo mv "$APP_NAME" "$INSTALL_PATH/"

echo "Cleaning up..."
rm "$FILE_NAME"

echo "Installed successfully! Run with:"
echo "  $APP_NAME --help"
