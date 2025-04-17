# fpm-cli-go

**fpm-cli-go** is a powerful and flexible CLI tool written in Go for managing Flutter projects efficiently and consistently.

Whether you're spinning up a new app, generating boilerplate, or building APKs with env-specific configs — `fpm` makes it seamless ⚡.

---

## 🛠 Features

- 🚀 Create new Flutter projects with Git, README, and state management boilerplate (GetX, BLoC, Provider, Riverpod)
- 🔧 Initialize Android signing configuration (`keytool`, `key.properties`)
- 📦 Build Flutter apps (APK or App Bundle) with environment and version info
- 📜 Auto-generate `README.md`, `.gitignore`, and Git setup
- 🐚 Shell completion scripts for Bash, Zsh, Fish, PowerShell
- 🛠 Modular Go structure using Cobra for future extension

---

## ✅ Completed Roadmap

- [x] `create` command: Generate Flutter app with boilerplate and Git
- [x] `build` command: Build APK/AAB for environments: dev/staging/prod
- [x] `signing` command: Generate Android keystore and properties file
- [x] Shell completions (Bash, Zsh, Fish, PowerShell)
- [x] Makefile for builds, tests, packaging
- [x] Downloadable binaries for Linux, macOS, Windows

---

## 🔮 Upcoming Features

- [ ] Pre-built templates (GetX, Clean Architecture)
- [ ] Auto environment file generation (`.env`)
- [ ] Git pre-hooks (format, lint before commit)
- [ ] Flutter pub commands (`clean`, `pub get`, etc.)
- [ ] IDE open command (`--vscode`, `--androidstudio`)

---

## 📦 Installation

### Option 1: Go Native

> Requires Go 1.18+

```bash
go install github.com/harsh6373/fpm-cli-go@latest
```

Then run:

```bash
fpm create
```

---

### Option 2: Download Precompiled Binaries

Download from the [Releases Page](https://github.com/harsh6373/fpm-cli-go/releases)

#### Example for Linux:

```bash
curl -L https://github.com/harsh6373/fpm-cli-go/releases/latest/download/fpm-linux -o /usr/local/bin/fpm
chmod +x /usr/local/bin/fpm
```

#### macOS or Windows:

Head to the GitHub [Releases](https://github.com/harsh6373/fpm-cli-go/releases) and grab your platform's binary.

---

### Option 3: Manual Clone + Build

```bash
git clone https://github.com/harsh6373/fpm-cli-go.git
cd fpm-cli-go
make install
```

---

## 🧪 Usage

```bash
fpm create     # Start a new Flutter project
fpm build      # Build APK/AAB with env + version
fpm signing    # Generate Android keystore & key.properties
```

---

## 🐚 Shell Completion

```bash
make completion
```

This generates completion scripts in the `completions/` folder for:
- Bash
- Zsh
- Fish
- PowerShell

---

## 🔨 Makefile Commands

```bash
make build         # Build for your OS
make test          # Run unit tests
make install       # Install globally via go install
make release       # Cross-compile for all platforms (Linux, macOS, Windows)
make completion    # Generate shell completion scripts
make clean         # Clean build artifacts
---
```

## 📁 Project Structure

```bash
fpm-cli-go/
├── cmd/            # Cobra commands: create, build, signing, etc.
├── utils/          # Reusable helpers: FS, Git, Signing, Readme
├── completions/    # Auto-generated shell completion scripts
├── bin/            # Compiled binaries go here (after `make build`)
├── main.go         # CLI entry point
├── go.mod
├── Makefile        # Build, test, cross-compile
```

---

## ❤️ Contributing

This is an open-source project built with ❤️ and Go. Contributions are **super welcome**!

Feel free to:
- Fork the repo
- Submit Pull Requests
- Open issues or suggest ideas

Let's build the best Flutter project manager CLI together!
