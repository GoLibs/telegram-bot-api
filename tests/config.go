package tests

type Tests struct {
	BotToken string
	ChatId   int64
	UserId   int64
}

func (t Tests) Defaults() Tests {
	return Tests{
		ChatId:   207260097,
		BotToken: "2014579763:AAECaWpB2_k_rxIw8xeiWG0m3lsw8q9oGzY",
		UserId:   81997375,
	}
}
