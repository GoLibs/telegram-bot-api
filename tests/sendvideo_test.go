package tests

import (
	"fmt"
	"testing"

	go_telegram_bot_api "github.com/golibs/telegram-bot-api"
)

func TestSendVideo(t *testing.T) {
	bot, err := go_telegram_bot_api.NewTelegramBot(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	message := (bot.Video()).SetChatId(Tests{}.Defaults().UserId).SetCaption("Love").SetVideoFilePath("video2.mkv")
	m, err := bot.Send(message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*m.Message.Video)

	message = (bot.Video()).SetChatId(Tests{}.Defaults().UserId).SetCaption("Love").SetVideoId("BAACAgQAAxkDAAECEwNeYWv3q6ZEzmcOUB8HQRwIlxMYYAACIgYAAiWdEVORmjxlkUFnfxgE")
	m, err = bot.Send(message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*m.Message.Video)

	/*m, err = (bot.Audio()).SetChatId(Tests{}.Defaults().UserId).SetCaption("Love").SetAudioThumbFilePath("photo.png").SetAudioId("CQACAgQAAxkDAAECEu5eYPQO1_MxAT-J8petsO2WrRRiTAACpwUAAiWdCVO8yR5RNMmjdBgE").Send()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*m.Message.Audio)*/
}
