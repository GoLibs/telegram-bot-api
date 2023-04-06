package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-telegram-bot-api"
)

func TestInit(t *testing.T) {
	bot, err := tgbotapi.New("638967213:AAEniQjX2XeOaGwouNU48G3v4Z0zIaHrz9.")
	if err != nil {
		fmt.Println(err)
		return
	}
	me := bot.GetMe()

	fmt.Println(me, "here 2")
}
