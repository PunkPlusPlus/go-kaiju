package main

import (
	"kaijuVpn/pkg/database"
	"kaijuVpn/pkg/telegram"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// connect to database
	database.DB.Connect()
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
		
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	telegramBot := telegram.NewBot(bot)
	err = telegramBot.Run()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
