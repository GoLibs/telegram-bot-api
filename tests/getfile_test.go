package tests

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/aliforever/go-telegram-bot-api"
)

func TestFile(t *testing.T) {
	bot, err := tgbotapi.New(Tests{}.Defaults().BotToken)
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
	fmt.Println(resp.File)
}

func TestFileDownload(t *testing.T) {
	bot, err := tgbotapi.New(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	up := (bot.File()).SetFileId("BQACAgQAAxkBAAII7WPulc5g6pWMDfn2VUEsEHuhUnMUAAJWDQAC0RJgU9B94T_-XLBKLgQ")

	resp, err := bot.SendWithOptions(up, tgbotapi.NewSendOptions().SetTimeout(time.Microsecond*1))
	if err != nil {
		fmt.Println(os.IsTimeout(err))
		fmt.Println(err)
		return
	}

	fmt.Println(*resp.File)
}
