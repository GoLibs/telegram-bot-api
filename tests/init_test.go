package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-telegram-bot-api"
)

func TestInit(t *testing.T) {
	bot, err := tgbotapi.NewTelegramBot("638967213:AAEniQjX2XeOaGwouNU48G3v4Z0zIaHrz9.")
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
