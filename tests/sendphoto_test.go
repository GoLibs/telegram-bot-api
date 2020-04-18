package tests

import (
	"fmt"
	"testing"

	go_telegram_bot_api "github.com/golibs/telegram-bot-api"
)

func TestSendPhoto(t *testing.T) {
	bot, err := go_telegram_bot_api.NewTelegramBot(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	//AgACAgQAAxkDAAECEudeYF9HGJ7VZzj6XzLdKNrNiy4IoAACDrIxGyWdCVNLOHwy3uUB1V0pthsABAEAAwIAA20AA1LzAwABGAQ
	p := bot.Photo()
	p.SetChatId(Tests{}.Defaults().UserId).SetPhotoFilePath("photo.png")
	m, err := bot.Send(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*m.Message.Photo)
	p = bot.Photo()
	p.SetChatId(Tests{}.Defaults().UserId).SetPhotoId("AgACAgQAAxkDAAECEwABXmFrWnTrgqevA4ChsqXoj0HaF7IAAhG0MRslnRFTNiO5tAeBc9FlP6gbAAQBAAMCAANtAAMZyQgAARgE")
	m, err = bot.Send(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*m.Message.Photo)
}
