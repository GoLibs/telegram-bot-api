package tests

import (
	"fmt"
	"testing"

	"github.com/kr/pretty"

	go_telegram_bot_api "github.com/GoLibs/telegram-bot-api"
)

func TestGetUserPhotos(t *testing.T) {
	bot, err := go_telegram_bot_api.NewTelegramBot(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	up := (bot.UserProfilePhotos()).SetUserId(Tests{}.Defaults().UserId)
	resp, err := bot.Send(up)
	if err != nil {
		fmt.Println(err)
		return
	}
	pretty.Println(resp.UserProfilePhotos)
}
