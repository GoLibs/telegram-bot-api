package tests

type Tests struct {
	BotToken string
	ChatId   int64
	UserId   int64
}

func (t Tests) Defaults() Tests {
	return Tests{
		ChatId:   ChatID,
		BotToken: BotToken,
		UserId:   UserID,
	}
}
