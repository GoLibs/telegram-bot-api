package structs

type ChatInviteLink struct {
	InviteLink  string `json:"invite_link"`
	Creator     *User  `json:"creator"`
	IsPrimary   bool   `json:"is_primary"`
	IsRevoked   bool   `json:"is_revoked"`
	ExpireDate  int64  `json:"expire_date"`
	MemberLimit int64  `json:"member_limit"`
}
