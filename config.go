package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// AlertConfig holds alert-specific settings
type AlertConfig struct {
	EnableTelegram   bool   `json:"enable_telegram"`
	TelegramBotToken string `json:"telegram_bot_token"`
	TelegramChatID   string `json:"telegram_chat_id"`
}

// Config holds the main TraceLock configuration
type Config struct {
	MonitorDir      string      `json:"monitor_dir"`
	LogFile         string      `json:"log_file"`
	ChangeThreshold int         `json:"change_threshold"`
	Alerts          AlertConfig `json:"alerts"`
}

// LoadConfig loads settings from config.json or applies defaults
func LoadConfig(path string) Config {
	var cfg Config

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("[!] config.json not found — using defaults.")
		cfg.MonitorDir = "./watched"
		cfg.LogFile = "./logs/tracelock.log"
		cfg.ChangeThreshold = 10
		return cfg
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		fmt.Println("[!] Invalid config.json — using defaults.")
		cfg.MonitorDir = "./watched"
		cfg.LogFile = "./logs/tracelock.log"
		cfg.ChangeThreshold = 10
	}

	return cfg
}
