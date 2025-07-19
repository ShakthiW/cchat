# Zsh plugin for cchat CLI
# Automatically downloads the correct binary from GitHub releases and adds it to your PATH

# Determine platform and architecture
get_os_arch() {
  case "$(uname -s)" in
    Darwin)
      OS="darwin"
      ;;
    Linux)
      OS="linux"
      ;;
    *)
      echo "❌ Unsupported OS: $(uname -s)"
      return 1
      ;;
  esac

  case "$(uname -m)" in
    x86_64)
      ARCH="amd64"
      ;;
    arm64|aarch64)
      ARCH="arm64"
      ;;
    *)
      echo "❌ Unsupported architecture: $(uname -m)"
      return 1
      ;;
  esac
}

install_cchat() {
  get_os_arch || return 1

  PLUGIN_DIR="${0:A:h}" # directory of this plugin file
  BIN_DIR="$PLUGIN_DIR/bin"
  CCHAT_BIN="$BIN_DIR/cchat"

  if [ ! -f "$CCHAT_BIN" ]; then
    echo "⬇️  Installing cchat for $OS-$ARCH..."
    mkdir -p "$BIN_DIR"
    URL="https://github.com/ShakthiW/chatcli/releases/latest/download/cchat-${OS}-${ARCH}"
    curl -sSL "$URL" -o "$CCHAT_BIN"
    chmod +x "$CCHAT_BIN"
    echo "✅ cchat installed successfully."
  fi

  export PATH="$BIN_DIR:$PATH"
}

install_cchat
