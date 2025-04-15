# fpm-cli-go

**fpm-cli-go** is a powerful and flexible CLI tool written in Go for managing Flutter projects more efficiently.

Whether you're spinning up a new app, automating common Flutter commands, or integrating Git setup out-of-the-box, `fpm` makes your workflow faster and smarter ⚡.

---

## 🛠 Features

- 🚀 Create new Flutter projects via `flutter create`
- ⚙️ Run common Flutter commands (build, clean, pub get, etc.)
- 📝 Automatically generate `.gitignore` and `README.md`
- 🔧 Initialize Git repo for new projects
- 🖥 Open projects in IDEs (VSCode, Android Studio)
- 📦 Scalable CLI structure using Cobra

---

## 🚀 TODO / Feature Roadmap

- [x] Create Flutter project with Git, README, and .gitignore
- [x] Run common Flutter commands (clean, build, pub get)
- [x] Open project in IDEs like VSCode and Android Studio

### 🔜 Upcoming Features

- [ ] Enhance `create` command:
  - Prompt user for:
    - Project name
    - Package name
    - Description
  - Ask user to select a state management approach:
    - GetX
    - BLoC
    - Provider
    - Riverpod
  - Add boilerplate code and dependencies based on selected option

- [ ] New `build` command:
  - Prompt user to:
    - Select build environment (dev, staging, prod)
    - Choose between APK or App Bundle
    - Enter version and build number
  - Generate final build file with a custom name based on inputs

- [ ] Scaffold GetX/BLoC folder structure and initial files
- [ ] Configurable project templates
- [ ] Auto-generate `.env` files for environments
- [ ] Optional pre-commit Git hooks for linting/formatting

## 📦 Installation

Clone the repo:

```bash
git clone https://github.com/yourusername/fpm-cli-go.git
cd fpm-cli-go
go build -o fpm

 
 ## 📁 Project structure

// fpm-cli-go/
// ├── cmd/
// │   ├── create.go
// │   └── build.go
// ├── helpers/
// │   ├── flutter.go
// │   └── fs.go
// ├── templates/
// │   └── ... (future use for templated files)
// ├── main.go
// └── go.mod

// ========================================
// 📄 main.go
