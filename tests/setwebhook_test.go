package tests

import (
	"fmt"
	"testing"

	go_telegram_bot_api "github.com/GoLibs/telegram-bot-api"
)

func TestSetWebhook(t *testing.T) {
	bot, err := go_telegram_bot_api.NewTelegramBot(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	bot.ListenWebhook("http://localhost:8082")
	res, err := bot.SetWebhook().SetIPAddress("127.0.0.1").Set()
	fmt.Println(res, err)
}
