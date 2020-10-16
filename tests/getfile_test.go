package tests

import (
	"fmt"
	"testing"

	"github.com/kr/pretty"

	go_telegram_bot_api "github.com/GoLibs/telegram-bot-api"
)

func TestFile(t *testing.T) {
	bot, err := go_telegram_bot_api.NewTelegramBot(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	up := (bot.File()).SetFileId("BQACAgQAAxkDAAECEwZeYXI771gPsruqrYEjhVewsKz8qAACJgYAAiWdEVPMjo7I4Zz4mxgE")
	resp, err := bot.Send(up)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = resp.File.Download("")
	if err != nil {
		fmt.Println(err)
		return
	}
	pretty.Println(resp.File)
}
