package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-telegram-bot-api"
)

func TestSendLocation(t *testing.T) {
	bot, err := tgbotapi.New(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	loc := (bot.Location()).SetChatId(Tests{}.Defaults().UserId).SetLatitude(20.073750).SetLongitude(30.280556).SetLivePeriod(60)
	resp, err := bot.Send(loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Message)
}
