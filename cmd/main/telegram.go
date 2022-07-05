package main

import (
	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"reminder/config"
)

var TIMEOUT = 60
var OFFSET = 0

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	config := config.NewConfig()

	bot, err := tgBotApi.NewBotAPI(config.TelegramConfig.Token)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	u := tgBotApi.NewUpdate(OFFSET)
	u.Timeout = TIMEOUT

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		repliedMessage := "hello"
		msg := tgBotApi.NewMessage(update.Message.Chat.ID, repliedMessage)
		bot.Send(msg)
	}
}
