# Go-Telegram-Bot-API

## Download
`go get -u github.com/aliforever/go-telegram-bot-api.git`

## Import
`import "github.com/aliforever/go-telegram-bot-api"`

## Usage
```go
botToken := "" // Place your bot token here
bot, err := tgbotapi.NewTelegramBot(botToken)
if err != nil {
    fmt.Println(err)
    return
}

chatId := int64(0) // Place your chatId here
bot.Send(bot.Message().SetChatId(chatId).SetText("Hello!"))
```