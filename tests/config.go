package tests

type Tests struct {
	BotToken string
	ChatId   int64
	UserId   int64
}

func (t Tests) Defaults() Tests {
	return Tests{
		ChatId:   207260097,
		BotToken: "1639631468:AAGMLl1184VfQRJF6lX_QClfrgg-HAxuRuA",
		UserId:   81997375,
	}
}
