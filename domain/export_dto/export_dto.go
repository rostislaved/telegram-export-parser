package export_dto

import (
	"telegram-export-parser/domain/message"
)

type ExportDTO struct {
	Name     string            `json:"name"`
	Type     string            `json:"type"`
	ID       int               `json:"id"`
	Messages []message.Message `json:"messages"`
}
