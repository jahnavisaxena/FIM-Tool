package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println(" TraceLock — Digital Forensic File Integrity Tool ")
	fmt.Println("----------------------------------------------------------")

	// 1️⃣ Load configuration
	cfg := LoadConfig("config.json")

	// 2️⃣ Ensure directories exist
	os.MkdirAll("logs", 0755)
	os.MkdirAll("reports", 0755)
	os.MkdirAll(cfg.MonitorDir, 0755)

	// 3️⃣ Setup log file
	logFile, err := os.OpenFile(cfg.LogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Cannot open log file:", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	// 4️⃣ Initialize external modules
	InitTelegram(cfg)             
	InitIntelligence(cfg.ChangeThreshold)

	// 5️⃣ Initialize baseline
	baselineFile := "baseline.json"
	if _, err := os.Stat(baselineFile); os.IsNotExist(err) {
		CreateBaseline(cfg.MonitorDir, baselineFile)
		SaveSignature(baselineFile)
	}

	// 6️⃣ Verify baseline integrity on startup
	ok, err := VerifySignature(baselineFile)
	if err != nil {
		log.Printf("[⚠️] Baseline signature missing: %v", err)
	} else if !ok {
		log.Printf("[🚨] Baseline integrity verification FAILED — possible tampering detected!")
		SendTelegramAlert("🚨 *TraceLock Critical Alert*\n\nBaseline file has been tampered with!")
	} else {
		log.Println("[✅] Baseline verified successfully.")
	}

	// 7️⃣ Load baseline
	baseline := LoadBaseline(baselineFile)

	// 8️⃣ Start monitoring
	go WatchDirectory(cfg, baseline, baselineFile)

	// 9️⃣ Graceful shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)
	<-done

	fmt.Println("\n🛑 Monitoring stopped.")
}
