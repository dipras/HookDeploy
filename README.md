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

## 📦 Setup
Make sure you have **Go 1.21+** installed.

### 1. Installing
```bash
git clone https://github.com/dipras/HookDeploy.git
cd hookdeploy
go mod tidy
```
### 2. Add .env file
You can check the example in `.env.example` file

### 3. Set your script
The script for deployment is using yaml named `config.yaml`, you can check the example in `config.yaml.example`

## 🌟 Next Feature
- [ ] Multi Repository
- [ ] Logging & Monitoring
- [ ] Web Dashboard
- [ ] Environment Variable
- [ ] Notification

## License
This project is licensed under the [GNU GPLv2](https://www.gnu.org/licenses/old-licenses/gpl-2.0.html).
