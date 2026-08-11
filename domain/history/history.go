package history

import (
	"telegram-export-parser/domain/message"
)

type History struct {
	Name     string            `json:"name"`
	Type     string            `json:"type"`
	ID       int               `json:"id"`
	Messages []message.Message `json:"messages"`
}
