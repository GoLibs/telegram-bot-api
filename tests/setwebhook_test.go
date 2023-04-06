package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-telegram-bot-api"
)

func TestSetWebhook(t *testing.T) {
	bot, err := tgbotapi.New(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	bot.ListenWebhook("http://localhost:8082")
	res, err := bot.SetWebhook().SetIPAddress("127.0.0.1").Set()
	fmt.Println(res, err)
}
