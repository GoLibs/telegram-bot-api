package tests

import (
	"fmt"
	"testing"

	go_telegram_bot_api "github.com/GoLibs/telegram-bot-api"
)

func TestSendPoll(t *testing.T) {
	bot, err := go_telegram_bot_api.NewTelegramBot(Tests{}.Defaults().BotToken)
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
