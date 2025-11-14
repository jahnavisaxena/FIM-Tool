package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type TelegramConfig struct {
	EnableTelegram   bool   `json:"enable_telegram"`
	TelegramBotToken string `json:"telegram_bot_token"`
	TelegramChatID   string `json:"telegram_chat_id"`
}

var globalTelegramConfig TelegramConfig

// Initialize Telegram settings from config.json
func InitTelegram(cfg Config) {
	globalTelegramConfig = TelegramConfig{
		EnableTelegram:   cfg.Alerts.EnableTelegram,
		TelegramBotToken: cfg.Alerts.TelegramBotToken,
		TelegramChatID:   cfg.Alerts.TelegramChatID,
	}
}

// Send alert to Telegram chat
func SendTelegramAlert(message string) {
	if !globalTelegramConfig.EnableTelegram {
		return
	}

	apiURL := fmt.Sprintf(
		"https://api.telegram.org/bot%s/sendMessage",
		globalTelegramConfig.TelegramBotToken,
	)

	payload := map[string]string{
		"chat_id":    globalTelegramConfig.TelegramChatID,
		"text":       message,
		"parse_mode": "Markdown", // Allows bold, italics, formatting
	}

	data, _ := json.Marshal(payload)

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Printf("[⚠️] Failed to send Telegram alert: %v", err)
		return
	}
	defer resp.Body.Close()

	log.Println("[📨] Telegram alert sent successfully.")
}
