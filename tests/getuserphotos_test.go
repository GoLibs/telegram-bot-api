package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-telegram-bot-api"
)

func TestGetUserPhotos(t *testing.T) {
	bot, err := tgbotapi.NewTelegramBot(Tests{}.Defaults().BotToken)
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
	fmt.Println(resp.UserProfilePhotos)
}
