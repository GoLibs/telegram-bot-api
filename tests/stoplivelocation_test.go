package tests

import (
	"fmt"
	"testing"

	go_telegram_bot_api "github.com/GoLibs/telegram-bot-api"
)

func TestStopLiveLocation(t *testing.T) {
	bot, err := go_telegram_bot_api.NewTelegramBot(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	loc := (bot.StopLiveLocation()).SetChatId(Tests{}.Defaults().UserId).SetMessageId(136532)
	resp, err := bot.Send(loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}
