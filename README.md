# 🛡️ FIMon — Real-Time File Integrity Monitoring Tool

**FIMon** is a lightweight, developer-friendly, real-time **File Integrity Monitoring (FIM)** tool written in Go.  
It continuously watches directories, computes **SHA-256 hashes**, and logs or alerts on any unauthorized file modifications.

---

##  Features

- 🔍 Real-time file monitoring using [`fsnotify`](https://github.com/fsnotify/fsnotify)
- 🧮 SHA-256 hash verification for integrity checks
- 🧾 Automatic baseline generation and updates
- 📜 Logging with timestamps and event type
- 🧠 Modular design for extension (email alerts, anomaly detection, DB storage)

---

##  Tech Stack

| Component | Tool / Package | Purpose |
|------------|----------------|----------|
| Language | Go (1.22+) | Core language |
| Monitoring | `fsnotify` | Real-time event watcher |
| Hashing | `crypto/sha256` | File integrity validation |
| Logging | Go `log` | Structured event logs |
| Storage | JSON baseline | Lightweight persistence |

---

## ⚙️ Installation

```bash
# Clone repository
git clone https://github.com/jahnavisaxena/FIMon.git
cd FIMon

# Install dependencies
go mod tidy

# Build executable
go build -o fimon
