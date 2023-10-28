package main

import (
	"os"
	"log"
	"strings"

	// ~ Paquetes locales
	"go-telegram-bot/functions"
	"go-telegram-bot/utils"

	// ~ Paquetes externos
	goDotEnv "github.com/joho/godotenv"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	goDotEnv.Load()
	log.Println("Starting bot...")

	bot, err := tgbotapi.NewBotAPI(os.Getenv("COWBOYBEBOP_TOKEN"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch strings.ToLower(update.Message.Command()) {
		case "help":
			msg.Text = functions.Help()
		case "hello":
			msg.Text = functions.Hello()
		case "status":
			msg.Text = functions.Status()
		case "items":
			// Obtener los productos desde functions.Items()
			products := functions.Items()

			// Enviar un mensaje para cada producto
			for _, product := range products {
				msg.Text = product
				if _, err := bot.Send(msg); err != nil {
					log.Println("Error al enviar mensaje:", err)
				}
			}

			// No continuar para evitar enviar un mensaje vacío después del bucle
			continue
		default:
			msg.Text = utils.Unknown()
		}

		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}

}
