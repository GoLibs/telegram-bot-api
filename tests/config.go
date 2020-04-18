package tests

type Tests struct {
	BotToken string
	ChatId   int64
	UserId   int64
}

func (t Tests) Defaults() Tests {
	return Tests{
		ChatId:   207260097,
		BotToken: "638967213:AAEniQjX2XeOaGwouNU48G3v4Z0zIaHrz9o",
		UserId:   81997375,
	}
}
