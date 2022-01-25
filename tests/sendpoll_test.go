package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-telegram-bot-api"
)

func TestSendPoll(t *testing.T) {
	bot, err := tgbotapi.NewTelegramBot(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	poll := (bot.Poll()).SetChatId(Tests{}.Defaults().UserId).SetQuestion("You There?").SetOptions([]string{"yes", "no"})
	resp, err := bot.Send(poll)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Message)
}
