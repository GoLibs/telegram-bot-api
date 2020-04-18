package tests

import (
	"fmt"
	"testing"

	go_telegram_bot_api "github.com/golibs/telegram-bot-api"
)

func TestSendSticker(t *testing.T) {
	bot, err := go_telegram_bot_api.NewTelegramBot(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	loc := (bot.Sticker()).SetChatId(Tests{}.Defaults().UserId).SetStickerFilePath("photo.png")
	resp, err := bot.Send(loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Message)
}
