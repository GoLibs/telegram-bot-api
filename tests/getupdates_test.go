package tests

import (
	"fmt"
	tgbotapi "github.com/aliforever/go-telegram-bot-api"
	"testing"
)

func TestGetUpdates(t *testing.T) {
	bot, err := tgbotapi.New(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		err := bot.GetUpdates().LongPoll()
		if err != nil {
			panic(err)
		}
	}()

	for update := range bot.Updates() {
		bot.Send(bot.Message().SetChatId(update.Message.Chat.Id).SetText("hello"))

		if update.Message.Video != nil {
			file := (bot.File()).SetFileId(update.Message.Video.FileId)
			res, err := bot.Send(file)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(res)
			return
		}

		return
	}
}
