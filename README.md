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

## 📦 Installation

Clone the repo:

```bash
git clone https://github.com/yourusername/fpm-cli-go.git
cd fpm-cli-go
go build -o fpm
