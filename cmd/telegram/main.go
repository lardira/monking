package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/lardira/monking/internal/bot/telegram"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}
}

func main() {
	token := os.Getenv("TELEGRAM_API_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_API_TOKEN env is required")
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancel()

	bot, err := telegram.New(token)
	if err != nil {
		log.Fatalf("could not bootstrap telegram bot: %v", err)
	}

	bot.Start(ctx)
	log.Println("Telegram bot shutdown")
}
