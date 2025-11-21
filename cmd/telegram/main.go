package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/lardira/monking/internal/bot/telegram"
	"github.com/lardira/monking/internal/env"
	"github.com/lardira/monking/internal/service"

	"github.com/lardira/monking/internal/db/sqlite"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}
}

func main() {
	db, err := sqlite.New(env.MustGetEnv("DATABASE_NAME"))
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}
	defer db.Close()

	token := env.MustGetEnv("TELEGRAM_API_TOKEN")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancel()

	userRepository := sqlite.NewUserRepository()

	userService := service.NewUserService(userRepository)

	bot, err := telegram.New(token, db, userService)
	if err != nil {
		log.Fatalf("could not bootstrap telegram bot: %v", err)
	}

	bot.Start(ctx)
	log.Println("Telegram bot shutdown")
}
