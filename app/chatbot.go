package app

import (
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func GetBot() *linebot.Client {

	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	return bot
}

func Broadcast(bot *linebot.Client, messageText string) {
	messageString := messageText
	message := linebot.NewTextMessage(messageString)
	resp, err := bot.BroadcastMessage(message).Do()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(*resp)
}
