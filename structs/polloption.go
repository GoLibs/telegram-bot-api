package structs

type PollOption struct {
	Text       string `json:"text"`
	VoterCount int64  `json:"voter_count"`
}
