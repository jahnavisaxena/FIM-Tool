# 🛡️ TraceLock — Development Roadmap

**TraceLock** is being developed as a modular, developer-friendly **Digital Forensic File Integrity Monitoring (FIM)** agent that evolves from a simple file monitor into a lightweight **Host-based Intrusion Detection System (HIDS/EDR)**.

---

## 🧭 Modular Development Phases

| **Phase** | **Name** | **Key Additions** | **Outcome** |
|:----------:|-----------|-------------------|--------------|
| 🧱 **1** | **Foundation Layer (Core FIM)** | Real-time file watcher, SHA-256 hash verification, baseline creation & JSON logging | Detects unauthorized file changes instantly and builds a reliable forensic trail |
| 🧩 **2** | **Intelligence Layer** | Email / Telegram alerts, anomaly detection, burst monitoring logic | Adds behavioral awareness and active alerting for suspicious or high-frequency modifications |
| 🗃️ **3** | **Persistence Layer** | SQLite forensic database, CLI queries, log rotation | Enables long-term forensic data retention and analyst queries |
| 🧠 **4** | **Awareness Layer** | Process & user context tracking (via OSQuery/Falco-like hooks), Sigma/YARA rule scanning | Correlates *who* changed *what*, moving TraceLock towards true HIDS capability |
| 🌐 **5** | **Visualization Layer** | Web dashboard (`net/http`), REST API, real-time analytics graphs | Converts TraceLock into an interactive forensic dashboard with visual monitoring |
| ⚙️ **6** | **Integration Layer** | SIEM/Wazuh/Elastic export, Docker containerization, systemd service | Enterprise-ready deployment with centralized SOC integration |

---

## 🧩 Version Progress Mapping

| **Version** | **Phase** | **Status** | **Key Features** |
|--------------|------------|------------|------------------|
| **v1.0** | 🧱 Foundation | ✅ Completed | Real-time monitoring, SHA-256 hash verification, logging |
| **v1.1** | 🧱 Foundation+ | ✅ Completed | Forensic evidence JSON logging (`forensic_log.json`) |
| **v1.2** | 🧱 Foundation+ | 🔜 Upcoming | Integrity verification for baseline & evidence files |
| **v2.0** | 🧩 Intelligence | ⏳ Planned | Alerts, anomaly detection, behavioral event tracking |
| **v2.5** | 🧩 Intelligence+ | ⏳ Planned | CLI command system with flags & subcommands |
| **v3.0** | 🗃️ Persistence | ⏳ Planned | SQLite forensic database with query and export support |
| **v3.5** | 🧠 Awareness | ⏳ Planned | Process/user correlation & YARA scanning |
| **v4.0** | 🌐 Visualization | ⏳ Future | Web dashboard, analytics, API-based visualization |
| **v5.0** | ⚙️ Integration | ⏳ Future | SIEM integration, Docker support, centralized logging |

---

## 🧠 Current Focus
We’re currently at **TraceLock v1.1**, which includes:
- 🔍 Real-time integrity monitoring  
- 🧮 SHA-256 file hashing  
- 🧾 Baseline creation and comparison  
- 📜 Forensic JSON evidence logging  

**Next milestone:**  
> 🧩 **v1.2 — Add baseline & forensic log integrity verification**  
Ensuring TraceLock’s own audit data cannot be tampered with.

---

## 🚀 Long-Term Vision
TraceLock’s ultimate goal is to become a **lightweight forensic EDR agent**, capable of:
- Real-time integrity detection  
- Evidence preservation  
- Alert automation  
- Forensic data visualization  
- SOC integration  

Built entirely in **Go**, TraceLock aims to stay **cross-platform**, **modular**, and **developer-centric** — enabling easy extension and integration in forensic labs, SOC pipelines, or research projects.

---

## 📅 Roadmap Summary

| **Milestone** | **Target Outcome** |
|----------------|--------------------|
| **v1.2** | Self-verifying forensic data (tamper-proof chain of custody) |
| **v2.0** | Real-time alerts & anomaly detection |
| **v2.5** | CLI interface for power users & automation |
| **v3.0** | SQLite forensic database |
| **v3.5** | Process/user context awareness |
| **v4.0+** | Web visualization & cloud readiness |

---

> 🧩 *TraceLock’s journey from a simple integrity monitor to a full forensic detection agent begins here.*  
> Contributions, ideas, and collaborations are always welcome.

---
