package tgbotapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/aliforever/go-telegram-bot-api/tools"

	"github.com/go-resty/resty/v2"
)

type TelegramBot struct {
	apiToken                    string
	apiUrl                      string
	username                    string
	recipientChatId             int64
	stopReceivingUpdatesChannel chan bool
	stoppedReceivingUpdates     bool
	updatesChannel              chan *Update
	client                      *resty.Client
	Tools                       tools.Tools
	logger                      resty.Logger
}

func NewTelegramBot(apiToken string) (tb *TelegramBot, err error) {
	rand.Seed(time.Now().UnixNano())
	address := fmt.Sprintf(`https://api.telegram.org/bot%s/`, apiToken)
	client := resty.New()
	client.SetHostURL(address)
	client.SetDoNotParseResponse(true)
	bot := &TelegramBot{
		apiToken:                    apiToken,
		apiUrl:                      address,
		stopReceivingUpdatesChannel: nil,
		client:                      client,
		Tools:                       tools.Tools{},
	}
	var resp *Response
	resp, err = bot.Send(bot.GetMe())
	if err != nil {
		return
	}
	if resp.User == nil {
		err = errors.New("empty_user_response")
	}
	tb = bot
	return
}

func NewTelegramBotWithLogger(apiToken string, logger resty.Logger) (tb *TelegramBot, err error) {
	rand.Seed(time.Now().UnixNano())
	address := fmt.Sprintf(`https://api.telegram.org/bot%s/`, apiToken)
	client := resty.New()
	client.SetHostURL(address)
	client.SetDoNotParseResponse(true)
	bot := &TelegramBot{
		apiToken:                    apiToken,
		apiUrl:                      address,
		stopReceivingUpdatesChannel: nil,
		client:                      client,
		Tools:                       tools.Tools{},
		logger:                      logger,
	}
	var resp *Response
	resp, err = bot.Send(bot.GetMe())
	if err != nil {
		return
	}
	if resp.User == nil {
		err = errors.New("empty_user_response")
	}
	tb = bot
	return
}

func (tb *TelegramBot) SetAPIServerUrl(address string) {
	if address[len(address)-1] == '/' {
		address = address[:len(address)-1]
	}
	tb.client.SetHostURL(fmt.Sprintf("%s/bot%s/", address, tb.apiToken))
}

func (tb *TelegramBot) SetRecipientChatId(chatId int64) {
	tb.recipientChatId = chatId
}

func (tb *TelegramBot) GetMe() (m *getMe) {
	m = &getMe{parent: tb}
	return
}

func (tb *TelegramBot) DeleteMessage() (m *deleteMessage) {
	m = &deleteMessage{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) Message() (m *sendMessage) {
	m = &sendMessage{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) CopyMessage() (m *copyMessage) {
	m = &copyMessage{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) EditMessageText() (m *editMessageText) {
	m = &editMessageText{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) EditMessageReplyMarkup() (m *editMessageReplyMarkup) {
	m = &editMessageReplyMarkup{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) ForwardMessage() (m *forwardMessage) {
	m = &forwardMessage{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) Photo() (m *sendPhoto) {
	m = &sendPhoto{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) Audio() (m *sendAudio) {
	m = &sendAudio{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) Video() (m *sendVideo) {
	m = &sendVideo{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) Voice() (m *sendVoice) {
	m = &sendVoice{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) VideoNote() (m *sendVideoNote) {
	m = &sendVideoNote{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) MediaGroup() (m *sendMediaGroup) {
	m = &sendMediaGroup{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) Document() (m *sendDocument) {
	m = &sendDocument{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) Animation() (m *sendAnimation) {
	m = &sendAnimation{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) Location() (m *sendLocation) {
	m = &sendLocation{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) Contact() (m *sendContact) {
	m = &sendContact{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) Venue() (m *sendVenue) {
	m = &sendVenue{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) Poll() (m *sendPoll) {
	m = &sendPoll{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) ChatAction() (m *sendChatAction) {
	m = &sendChatAction{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) UserProfilePhotos() (m *getUserProfilePhotos) {
	m = &getUserProfilePhotos{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetUserId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) File() (m *getFile) {
	m = &getFile{parent: tb}
	return
}

func (tb *TelegramBot) KickChatMember() (m *kickChatMember) {
	m = &kickChatMember{parent: tb}
	return
}

func (tb *TelegramBot) AnswerCallbackQuery() (m *answerCallbackQuery) {
	m = &answerCallbackQuery{parent: tb}
	return
}

func (tb *TelegramBot) PromoteChatMember() (m *promoteChatMember) {
	m = &promoteChatMember{parent: tb}
	return
}

func (tb *TelegramBot) RestrictChatMember() (m *restrictChatMember) {
	m = &restrictChatMember{parent: tb}
	return
}

func (tb *TelegramBot) UnbanChatMember() (m *unbanChatMember) {
	m = &unbanChatMember{parent: tb}
	return
}

func (tb *TelegramBot) SetChatAdministratorCustomTitle() (m *setChatAdministratorCustomTitle) {
	m = &setChatAdministratorCustomTitle{parent: tb}
	return
}

func (tb *TelegramBot) EditLiveLocation() (m *editMessageLiveLocation) {
	m = &editMessageLiveLocation{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) StopLiveLocation() (m *stopMessageLiveLocation) {
	m = &stopMessageLiveLocation{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) EditMessageMedia() (m *editMessageMedia) {
	m = &editMessageMedia{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) Sticker() (m *sendSticker) {
	m = &sendSticker{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) StickerFile() (m *uploadStickerFile) {
	m = &uploadStickerFile{parent: tb}
	return
}

func (tb *TelegramBot) GetUpdates() (m *getUpdates) {
	m = &getUpdates{parent: tb}
	return
}

func (tb *TelegramBot) SetWebhook() (m *setWebhook) {
	m = &setWebhook{parent: tb}
	return
}

func (tb *TelegramBot) DeleteWebhook() (m *deleteWebhook) {
	m = &deleteWebhook{parent: tb}
	return
}

func (tb *TelegramBot) StopReceivingUpdates() {
	close(tb.stopReceivingUpdatesChannel)
}

func (tb *TelegramBot) ListenWebhook(address string) (err error) {
	tb.updatesChannel = make(chan *Update)
	http.HandleFunc(fmt.Sprintf("/%s", tb.apiToken), func(writer http.ResponseWriter, request *http.Request) {
		defer request.Body.Close()
		var u *Update
		j := json.NewDecoder(request.Body)
		err := j.Decode(&u)
		if err != nil {
			// commit
			fmt.Println("error decoding update", err.Error())
			return
		}
		tb.updatesChannel <- u
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("ok"))
	})
	err = http.ListenAndServe(address, nil)
	return
}

func (tb *TelegramBot) Updates() chan *Update {
	// commit
	return tb.updatesChannel
}

func (tb *TelegramBot) Send(config Config) (result *Response, err error) {
	request := tb.client.R()

	err = tb.prepareRequest(config, request)
	if err != nil {
		return nil, err
	}

	res, err := request.Execute(config.method(), config.endpoint())
	if err != nil {
		return nil, err
	}

	var raw []byte
	result, raw, err = tb.getMessageResponse(res, config)
	if err != nil {
		tb.logErrBytes(raw, err)
		return nil, err
	}

	return
}

func (tb *TelegramBot) SendWithOptions(config Config, options *SendOptions) (result *Response, err error) {
	client := tb.client

	if options != nil {
		if options.timeout != 0 {
			newClient := *client
			client = newClient.SetTimeout(options.timeout)
		}
	}

	request := client.R()

	err = tb.prepareRequest(config, request)
	if err != nil {
		return nil, err
	}

	res, err := request.Execute(config.method(), config.endpoint())
	if err != nil {
		return nil, err
	}

	var raw []byte
	result, raw, err = tb.getMessageResponse(res, config)
	if err != nil {
		tb.logErrBytes(raw, err)
		return nil, err
	}

	return
}

func (tb *TelegramBot) logErrBytes(rawBytes []byte, err error) {
	if tb.logger != nil {
		tb.logger.Errorf("%s : %s", err, string(rawBytes))
	}
}
