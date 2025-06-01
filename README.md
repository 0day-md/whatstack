## 🧠 WhatStack

**WhatStack** is a lightweight command-line tool that identifies the technologies used by a website. It uses the [Wappalyzer Go](https://github.com/projectdiscovery/wappalyzergo) library to fingerprint frameworks, CMSs, web servers, JavaScript libraries, and more — based on HTTP headers and HTML content.

> 🔍 Ideal for developers, bug bounty hunters, and researchers who want quick insight into a site's tech stack.

---

### 🚀 Features

* Detects web technologies with high accuracy
* Skips sites with bot protections (e.g., Cloudflare)
* Simple, fast, no unnecessary bloat
* Optionally save results to a file

---

### 💡 Inspired By

This tool is **inspired by** and **powered by**:

> 🔗 [`github.com/projectdiscovery/wappalyzergo`](https://github.com/projectdiscovery/wappalyzergo)
> A Go implementation of Wappalyzer’s fingerprinting logic.

Big thanks to the ProjectDiscovery team!

---

### 🛠️ Installation

1. **Clone the repo**:

```bash
git clone https://github.com/yourusername/whatstack.git
cd whatstack
```

2. **Build the binary**:

```bash
go build -o whatstack
```

> Requires [Go 1.18+](https://golang.org/dl/)

---

### 📦 Usage

```bash
./whatstack -d example.com
```

#### Optional Flags:

| Flag | Description                     |
| ---- | ------------------------------- |
| `-d` | Domain to analyze (required)    |
| `-o` | Output file to save the results |

#### Example:

```bash
./whatstack -d wikipedia.org -o stack.txt
```

---

### 🧱 Sample Output

```
🔍 Technologies detected:
- MediaWiki
- PHP
- jQuery
✅ Results saved to stack.txt
```

---

### ⚠️ Notes

* Sites with bot protections (e.g., Cloudflare, DDoS-Guard, captchas) are automatically skipped.
* Requests time out after 10 seconds to avoid hangs or slow responses.

---

### 📄 License

MIT License — see [`LICENSE`](LICENSE) for details.

---

### 👨‍💻 Author

Made with 🛠️ by **0day.md**
GitHub: [@0day-md](https://github.com/0day-md)

---

Let me know if you want a version with badges or GitHub Actions support!
