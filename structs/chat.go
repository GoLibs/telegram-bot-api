package structs

type Chat struct {
	Id               int64           `json:"id"`
	Type             string          `json:"type"`
	Title            string          `json:"title"`
	Username         string          `json:"username"`
	FirstName        string          `json:"first_name"`
	LastName         string          `json:"last_name"`
	Photo            ChatPhoto       `json:"photo"`
	Description      string          `json:"description"`
	InviteLink       string          `json:"invite_link"`
	PinnedMessage    *Message        `json:"pinned_message"`
	Permissions      ChatPermissions `json:"permissions"`
	SlowModeDelay    int64           `json:"slow_mode_delay"`
	StickerSetName   string          `json:"sticker_set_name"`
	CanSetStickerSet bool            `json:"can_set_sticker_set"`
}
