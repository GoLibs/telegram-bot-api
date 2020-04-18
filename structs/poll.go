package structs

type Poll struct {
	Id                    string       `json:"id"`
	Question              string       `json:"question"`
	Options               []PollOption `json:"options"`
	TotalVoterCount       int64        `json:"total_voter_count"`
	IsClosed              bool         `json:"is_closed"`
	IsAnonymous           bool         `json:"is_anonymous"`
	Type                  string       `json:"type"`
	AllowsMultipleAnswers bool         `json:"allows_multiple_answers"`
	CorrectOptionId       int64        `json:"correct_option_id"`
}
