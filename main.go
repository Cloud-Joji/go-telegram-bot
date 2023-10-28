package main

import (
	"os"
	"log"
	help "gitlab.com/CloudyJoji/go-telegram-bot/functions/help"
	goDotEnv "github.com/joho/godotenv"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	goDotEnv.Load()

	bot, err := tgbotapi.NewBotAPI(os.Getenv("COWBOYBEBOP_TOKEN"))
	if err != nil {
	    panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		// msg.ReplyToMessageID = update.Message.MessageID

		switch update.Message.Command() {
		case "help":
				msg.Text = help()
		case "hola":
				msg.Text = "Hey! :)"
		case "estado":
				msg.Text = "Fumando!"
		default:
				msg.Text = "???????"
		}


		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}

}
