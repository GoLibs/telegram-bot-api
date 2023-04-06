package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-telegram-bot-api"
)

func TestStopLiveLocation(t *testing.T) {
	bot, err := tgbotapi.New(Tests{}.Defaults().BotToken)
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
