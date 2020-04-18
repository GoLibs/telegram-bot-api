package tests

import (
	"fmt"
	"testing"

	go_telegram_bot_api "github.com/golibs/telegram-bot-api"
)

func TestUploadStickerFile(t *testing.T) {
	bot, err := go_telegram_bot_api.NewTelegramBot(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	//AgACAgQAAxkDAAECEudeYF9HGJ7VZzj6XzLdKNrNiy4IoAACDrIxGyWdCVNLOHwy3uUB1V0pthsABAEAAwIAA20AA1LzAwABGAQ
	p := bot.StickerFile()
	p.SetUserId(Tests{}.Defaults().UserId).SetStickerFilePath("photo.png")
	m, err := bot.Send(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*m.File)

}
