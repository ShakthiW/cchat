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

BINARY_URL="https://github.com/ShakthiW/chatcli/releases/latest/download/cchat-${OS}-${ARCH}"
INSTALL_DIR="$HOME/.local/bin"

mkdir -p "$INSTALL_DIR"

echo "⬇️  Installing cchat for $OS-$ARCH..."

curl -fsSL "$BINARY_URL" -o "$INSTALL_DIR/cchat"

chmod +x "$INSTALL_DIR/cchat"

if ! grep -q "$INSTALL_DIR" <<< "$PATH"; then
  echo "⚠️  Add $INSTALL_DIR to your PATH if not already present"
fi

# Add alias
if ! grep -q "alias ai=" ~/.zshrc; then
  echo "\n# Alias for cchat CLI" >> ~/.zshrc
  echo "alias ai='cchat chat'" >> ~/.zshrc
fi

echo "✅ cchat installed successfully!"
echo "⚙️  Restart your terminal or run 'source ~/.zshrc' to use the 'ai' alias."
