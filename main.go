package main

import (
	"fmt"
	"os"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	fmt.Println("Start main function")
	//// Get bot token from environment variable or command line argument
	token := os.Getenv("TELEGRAM_API_TOKEN")
	if token == "" {
		token = os.Args[1]
	}

	if token == "" {
		fmt.Println("Error: No token provided")
		return
	}

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		fmt.Println("Error while creating bot")
		return
	}

	bot.Debug = true
	fmt.Println("Bot created")
	fmt.Println("Bot name: ", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		fmt.Println("Error while getting updates")
		return
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		fmt.Println("Message: ", update.Message.Text)
		fmt.Println("From: ", update.Message.From.UserName)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
