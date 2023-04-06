package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-telegram-bot-api"
)

func TestEditLiveLocation(t *testing.T) {
	bot, err := tgbotapi.New(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	loc := (bot.EditLiveLocation()).SetChatId(Tests{}.Defaults().UserId).SetMessageId(136532).SetLatitude(21.073750).SetLongitude(23.280556)
	resp, err := bot.Send(loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}
