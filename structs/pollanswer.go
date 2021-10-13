package structs

type PollAnswer struct {
	PollId    string  `json:"poll_id"`
	User      *User   `json:"user"`
	OptionIds []int64 `json:"option_ids"`
}
