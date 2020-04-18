package structs

type UpdatesChannel <-chan Update

func (uc UpdatesChannel) Clear() {
	for len(uc) != 0 {
		<-uc
	}
}

type Update struct {
	UpdateId          int64    `json:"update_id"`
	Message           *Message `json:"message"`
	EditedMessage     *Message `json:"edited_message"`
	ChannelPost       *Message `json:"channel_post"`
	EditedChannelPost *Message `json:"edited_channel_post"`
	// InlineQuery        InlineQuery        `json:"inline_query"`
	// ChosenInlineResult ChosenInlineResult `json:"chosen_inline_result"`
	CallbackQuery *CallbackQuery `json:"callback_query"`
	// ShippingQuery      ShippingQuery      `json:"shipping_query"`
	// PreCheckoutQuery   PreCheckoutQuery   `json:"pre_checkout_query"`
	Poll Poll `json:"poll"`
	// PollAnswer         PollAnswer         `json:"poll_answer"`
}
