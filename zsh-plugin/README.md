# 🤖 cchat Oh My Zsh Plugin

This plugin enables you to use the `cchat` CLI tool directly from your terminal and adds a convenient alias `ai` for quick access.

## 🚀 Features

- Automatically installs the correct `cchat` binary based on your OS and architecture.
- Adds the binary to your terminal `PATH`.
- Sets up `ai` as an alias for `cchat chat`.
- Works seamlessly with **Oh My Zsh**.

---

## 📦 Installation

### 🔁 Option 1: Manual Installation

1. Clone the repo or download the plugin folder:

```sh
git clone https://github.com/ShakthiW/chatcli.git
```

2. Move into the plugin directory:

```sh
cd chatcli/zsh-plugin
```

3. Create a symlink to use as an Oh My Zsh plugin:

```sh
ln -s $(pwd) ~/.oh-my-zsh/custom/plugins/cchat
```

4. Enable the plugin by adding `cchat` to your plugin list in `~/.zshrc`:

```zsh
plugins=(git cchat)
```

5. Reload your shell:

```sh
source ~/.zshrc
```

Done! Now you can use:

```sh
ai "What's the weather like?"
```

---

### 🧪 Option 2: Install via Script

```sh
sh -c "$(curl -fsSL https://raw.githubusercontent.com/ShakthiW/chatcli/main/zsh-plugin/cchat.plugin.zsh)"
```

> 💡 This will download the correct binary and add the alias `ai` to your `.zshrc`.

---

## 🧩 Alias

| Command | Description                |
| ------- | -------------------------- |
| `ai`    | Shorthand for `cchat chat` |
| `cchat` | The base CLI tool binary   |

---

## 🛠️ Requirements

* Zsh shell
* [Oh My Zsh](https://ohmyz.sh/)
* curl

---

## 📁 File Structure

```
zsh-plugin/
├── cchat.plugin.zsh   # Plugin logic
└── README.md          # This file
```

---

## 🔐 Permissions

Make sure your plugin file is executable:

```sh
chmod +x zsh-plugin/cchat.plugin.zsh
```

---

## 📬 Issues or Feedback?

Open an issue at: [github.com/ShakthiW/chatcli/issues](https://github.com/ShakthiW/chatcli/issues)
