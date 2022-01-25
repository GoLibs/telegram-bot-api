package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-telegram-bot-api"
)

func TestForwardMessage(t *testing.T) {
	bot, err := tgbotapi.NewTelegramBot(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	fm := (bot.ForwardMessage()).SetChatId(Tests{}.Defaults().UserId).SetFromChatId(Tests{}.Defaults().UserId).SetMessageId(135905)

	m, err := bot.Send(fm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(m)

}
