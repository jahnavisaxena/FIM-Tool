package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var (
	eventCount     int
	eventLock      sync.Mutex
	lastResetTime  = time.Now()
	alertThreshold = 10
)

// InitIntelligence loads anomaly detection threshold from config.json
func InitIntelligence(threshold int) {
	alertThreshold = threshold
	go resetCounter()
}

// resetCounter clears event count every minute
func resetCounter() {
	for {
		time.Sleep(60 * time.Second)
		eventLock.Lock()
		eventCount = 0
		lastResetTime = time.Now()
		eventLock.Unlock()
	}
}

// TrackEvent is called for every file create/modify/delete event
func TrackEvent(eventType, filePath string) {
	eventLock.Lock()
	defer eventLock.Unlock()

	eventCount++

	// If threshold exceeded → anomaly detected
	if eventCount > alertThreshold {
		log.Printf("🚨 Anomaly detected! %d file events in the last minute — possible ransomware activity!", eventCount)

		// Create Telegram message
		alertMsg := fmt.Sprintf(
			"*🚨 TraceLock Alert: Suspicious Activity Detected!*\n\n"+
				"*%d file events* detected within one minute.\n"+
				"Possible *ransomware*, *mass file modification*, or *system compromise*.\n\n"+
				"*Timestamp:* %s",
			eventCount,
			time.Now().Format("2006-01-02 15:04:05"),
		)

		// Send alert
		SendTelegramAlert(alertMsg)

		// Reset event counter after alert
		eventCount = 0
	}
}
