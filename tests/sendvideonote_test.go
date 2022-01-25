package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-telegram-bot-api"
)

func TestSendVideoNote(t *testing.T) {
	bot, err := tgbotapi.NewTelegramBot(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	message := (bot.VideoNote()).SetChatId(Tests{}.Defaults().UserId).SetCaption("Love").SetVideoNoteFilePath("video.mkv")
	m, err := bot.Send(message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*m.Message.VideoNote)

	message = (bot.VideoNote()).SetChatId(Tests{}.Defaults().UserId).SetCaption("Love").SetVideoNoteId("DQACAgQAAxkDAAECEyJeYaMzAbcxbX80TGvh_mC4AdKCawACNwYAAiWdEVODQvIVVZoJJRgE")
	m, err = bot.Send(message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*m.Message.VideoNote)

	/*m, err = (bot.Audio()).SetChatId(Tests{}.Defaults().UserId).SetCaption("Love").SetAudioThumbFilePath("photo.png").SetAudioId("CQACAgQAAxkDAAECEu5eYPQO1_MxAT-J8petsO2WrRRiTAACpwUAAiWdCVO8yR5RNMmjdBgE").Send()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*m.Message.Audio)*/
}
