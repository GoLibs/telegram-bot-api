package responses

import (
	"fmt"
)

type Error struct {
	ErrorCode   int64  `json:"error_code,omitempty"`
	Description string `json:"description,omitempty"`
}

func (e Error) Error() string {
	return fmt.Sprintf(`%d - %s`, e.ErrorCode, e.Description)
}
