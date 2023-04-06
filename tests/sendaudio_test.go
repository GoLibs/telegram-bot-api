package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-telegram-bot-api"
)

func TestSendAudio(t *testing.T) {
	bot, err := tgbotapi.New(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	message := (bot.Audio()).SetChatId(Tests{}.Defaults().UserId).SetCaption("Love").SetAudioFilePath("beep.mp3")
	m, err := bot.Send(message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*m.Message.Audio)

	message = (bot.Audio()).SetChatId(Tests{}.Defaults().UserId).SetCaption("Love").SetAudioId("CQACAgQAAxkDAAECEu5eYPQO1_MxAT-J8petsO2WrRRiTAACpwUAAiWdCVO8yR5RNMmjdBgE")
	m, err = bot.Send(message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*m.Message.Audio)

	/*	m, err = (bot.Audio()).SetChatId(Tests{}.Defaults().UserId).SetCaption("Love").SetAudioThumbFilePath("photo.png").SetAudioId("CQACAgQAAxkDAAECEu5eYPQO1_MxAT-J8petsO2WrRRiTAACpwUAAiWdCVO8yR5RNMmjdBgE").Send()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(*m.Message.Audio)*/
}
