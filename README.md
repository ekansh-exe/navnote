<div align="center">

```
███╗   ██╗ █████╗ ██╗   ██╗███╗   ██╗ ██████╗ ████████╗███████╗
████╗  ██║██╔══██╗██║   ██║████╗  ██║██╔═══██╗╚══██╔══╝██╔════╝
██╔██╗ ██║███████║██║   ██║██╔██╗ ██║██║   ██║   ██║   █████╗  
██║╚██╗██║██╔══██║╚██╗ ██╔╝██║╚██╗██║██║   ██║   ██║   ██╔══╝  
██║ ╚████║██║  ██║ ╚████╔╝ ██║ ╚████║╚██████╔╝   ██║   ███████╗
╚═╝  ╚═══╝╚═╝  ╚═╝  ╚═══╝  ╚═╝  ╚═══╝ ╚═════╝    ╚═╝   ╚══════╝
```

**A minimal, keyboard-driven note-taking app that lives in your terminal.**  
Notes are plain `.md` files. No database. No cloud. No nonsense.

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go&logoColor=white)
![License](https://img.shields.io/badge/license-MIT-blueviolet?style=flat-square)
![Built with Bubble Tea](https://img.shields.io/badge/built%20with-Bubble%20Tea-EE6FF8?style=flat-square)

</div>

---

## What is NavNote?

NavNote is a TUI (terminal user interface) note-taking app written in Go. It stores everything as plain Markdown files in `~/.navnote/` — so your notes are always yours, readable by any editor, and easy to sync with Git, Dropbox, or whatever you prefer.

No accounts. No subscriptions. No Electron. Just your terminal and your thoughts.

> 🌿Just like Bulbasaur is the best. NavNote is the best terminal notes app.

---

## Features

- 📝 **Create & edit** Markdown notes without leaving the terminal
- 📋 **Browse** all notes in a filterable, searchable list
- 💾 **Save** with a single keypress
- 🎨 **Lavender-themed TUI** built with [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- 📁 **Plain `.md` files** — no lock-in, works with any editor
- ⚡ **Fast** — no startup time, no loading screens
- 🔍 **Fuzzy filter** notes by name in the list view

---

## Installation

> ⚠️ **Make sure you have Go 1.21+ installed first** — https://go.dev/doc/install

### Option 1 — `go install` (recommended, no cloning needed)

```bash
go install github.com/ekansh-exe/navnote@latest
```

Make sure your Go binary path is on `$PATH`:

```bash
# Add this to your ~/.bashrc or ~/.zshrc if not already there
export PATH="$PATH:$(go env GOPATH)/bin"
```

Then just run:

```bash
navnote
```

---

### Option 2 — Build from source

```bash
git clone https://github.com/ekansh-exe/navnote.git
cd navnote
go build -o navnote .
./navnote
```

---

### Option 3 — Using Make

```bash
git clone https://github.com/ekansh-exe/navnote.git
cd navnote
make run        # run directly without building
make build      # build the binary
make install    # install globally via go install
```

---

## Requirements

- **Go 1.21+** — [Install Go](https://go.dev/doc/install)
- A terminal with **256 color support** (any modern terminal: iTerm2, Alacritty, WezTerm, GNOME Terminal, Windows Terminal, etc.)

---

## Keyboard Shortcuts

| Key | Action |
|-----|--------|
| `Ctrl+N` | Create a new note |
| `Ctrl+L` | Open the notes list |
| `Enter` | Open the selected note |
| `Ctrl+S` | Save the current note |
| `Esc` | Go back / cancel current action |
| `Ctrl+C` or `q` | Quit |
| `Ctrl+D` | Permanently delete the selected file after confirmation|

> **Tip:** In the list view, just start typing to fuzzy-filter your notes by name.

---

## Project Structure

```
navnote/
├── main.go       # Entry point — wires up the Bubble Tea program
├── model.go      # TUI model: struct, Init, Update, View
├── files.go      # Vault logic: init, listFiles, item type
├── styles.go     # All lipgloss styles and color variables
├── Makefile      # Build, run, install, clean targets
├── go.mod
├── go.sum
└── README.md
```

All files share `package main` — Go builds the whole directory together, so they freely reference each other's types and functions with no internal imports needed.

---

## Where are my notes stored?

```
~/.navnote/
├── meeting-notes.md
├── ideas.md
└── todo.md
```

Plain Markdown. Open them in Neovim, VS Code, Obsidian — whatever you want. NavNote doesn't care.

---

## Contributing

Issues and PRs are welcome. For big changes, open an issue first so we can talk about it.

```bash
git clone https://github.com/ekansh-exe/navnote.git
cd navnote
make run   # start hacking
```

---

## License

MIT — do whatever you want with it.

---

<div align="center">
  Made with 🌿 and a strong opinion about Bulbasaur.
</div>