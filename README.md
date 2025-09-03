# HookDeploy 🚀

**HookDeploy** is a lightweight **Go** application designed to automate deployments using **GitHub Webhooks**.  
With HookDeploy, every time code changes (push) occur in your repository, it can execute a sequence of commands defined in your configuration.

---

## ✨ Features
- 📡 **Receive GitHub Webhooks** → automatically detects `push` events.
- ⚙️ **Custom Command Execution** → run commands defined in `config.yaml`.
- 🛡️ **Signature Verification** → ensure the request is truly from GitHub (HMAC SHA-256).
- 📂 **Simple Configuration** → just a YAML file, easy to read and edit.
- 🔧 **Single Repository Focus** → lightweight and simple, perfect for personal projects or small servers.

---

## 📦 Installation
Make sure you have **Go 1.21+** installed.

```bash
git clone https://github.com/username/hookdeploy.git
cd hookdeploy
go mod tidy
go run main.go
```

## License
This project is licensed under the [GNU GPLv2](https://www.gnu.org/licenses/old-licenses/gpl-2.0.html).