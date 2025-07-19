#!/bin/zsh

set -e

ARCH=$(uname -m)
OS=$(uname | tr '[:upper:]' '[:lower:]')

if [[ "$ARCH" == "arm64" ]]; then
  ARCH="arm64"
elif [[ "$ARCH" == "x86_64" ]]; then
  ARCH="amd64"
else
  echo "❌ Unsupported architecture: $ARCH"
  return 1
fi

INSTALL_DIR="$HOME/.local/bin"
TMP_DIR=$(mktemp -d)

# Construct the tar.gz URL
TAR_URL="https://github.com/ShakthiW/chatcli/releases/latest/download/cchat-${OS}-${ARCH}.tar.gz"

echo "⬇️  Downloading cchat from $TAR_URL..."
mkdir -p "$INSTALL_DIR"
curl -fsSL "$TAR_URL" -o "$TMP_DIR/cchat.tar.gz"

echo "📦 Extracting..."
tar -xzf "$TMP_DIR/cchat.tar.gz" -C "$TMP_DIR"

# Move binary to install dir
mv "$TMP_DIR/cchat" "$INSTALL_DIR/cchat"
chmod +x "$INSTALL_DIR/cchat"

# Clean up
rm -rf "$TMP_DIR"

# Add to PATH if not already
if ! echo "$PATH" | grep -q "$INSTALL_DIR"; then
  echo "\n# Add cchat to PATH" >> ~/.zshrc
  echo "export PATH=\"\$HOME/.local/bin:\$PATH\"" >> ~/.zshrc
fi

# Add alias for ease of use
if ! grep -q "alias ai=" ~/.zshrc; then
  echo "\n# Alias for cchat CLI" >> ~/.zshrc
  echo "alias ai='cchat chat'" >> ~/.zshrc
fi

echo "✅ cchat installed successfully!"
echo "⚙️  Restart your terminal or run 'source ~/.zshrc' to use the 'ai' alias."
