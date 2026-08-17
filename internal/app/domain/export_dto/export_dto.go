package export_dto

import (
	"telegram-export-parser/internal/app/domain/message"
)

type ExportDTO struct {
	ID       int               `json:"id"`
	Name     string            `json:"name"`
	Type     string            `json:"type"`
	Messages []message.Message `json:"messages"`
}
