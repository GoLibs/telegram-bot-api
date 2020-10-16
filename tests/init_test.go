package tests

import (
	"fmt"
	"testing"

	go_telegram_bot_api "github.com/GoLibs/telegram-bot-api"
)

func TestInit(t *testing.T) {
	bot, err := go_telegram_bot_api.NewTelegramBot("638967213:AAEniQjX2XeOaGwouNU48G3v4Z0zIaHrz9.")
	if err != nil {
		fmt.Println(err)
		return
	}
	me, err := bot.GetMe()
	if err != nil {
		fmt.Println(err, "here")
		return
	}
	fmt.Println(me, "here 2")
}
