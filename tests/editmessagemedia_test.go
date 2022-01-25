package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-telegram-bot-api"
)

func TestEditMessageMedia(t *testing.T) {
	bot, err := tgbotapi.NewTelegramBot(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	message := (bot.Photo()).SetChatId(Tests{}.Defaults().UserId).SetCaption("Love").SetPhotoFilePath("photo.png")
	m, err := bot.Send(message)
	if err != nil {
		fmt.Println(err, "here")
		return
	}

	edit := (bot.EditMessageMedia()).SetChatId(Tests{}.Defaults().UserId).SetMessageId(m.Message.MessageId)
	(edit.InputPhoto()).SetPhotoPath("photo2.png").SetCaption("Hate")
	m, err = bot.Send(edit)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*m.Message.Photo)

	messageVideo := (bot.Video()).SetChatId(Tests{}.Defaults().UserId).SetCaption("Love").SetVideoFilePath("video.mkv")
	m, err = bot.Send(messageVideo)
	if err != nil {
		fmt.Println(err, "here")
		return
	}
	edit = (bot.EditMessageMedia()).SetChatId(Tests{}.Defaults().UserId).SetMessageId(m.Message.MessageId)
	(edit.InputVideo()).SetVideoPath("video2.mkv").SetCaption("Hate")
	m, err = bot.Send(edit)
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
