package tests

import (
	"fmt"
	"testing"

	go_telegram_bot_api "github.com/golibs/telegram-bot-api"
)

func TestForwardMessage(t *testing.T) {
	bot, err := go_telegram_bot_api.NewTelegramBot(Tests{}.Defaults().BotToken)
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
