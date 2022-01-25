# Go-Telegram-Bot-API

## Download
`go get -u github.com/aliforever/go-telegram-bot-api.git`

## Usage
```go
bot, err := tgbotapi.NewTelegramBot(Tests{}.Defaults().BotToken)
if err != nil {
    fmt.Println(err)
    return
}
```