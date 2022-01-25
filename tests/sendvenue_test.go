package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-telegram-bot-api"
)

func TestSendVenue(t *testing.T) {
	bot, err := tgbotapi.NewTelegramBot(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	loc := (bot.Venue()).SetChatId(Tests{}.Defaults().UserId).SetLatitude(20.073750).SetLongitude(30.280556)
	resp, err := bot.Send(loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Message)
}
