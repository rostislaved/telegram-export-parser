package text_entity

import (
	"encoding/json"
	"fmt"
	"log/slog"
)

type TextEntity struct {
	Type       Type   `json:"type"`
	Text       string `json:"text"`
	Collapsed  bool   `json:"collapsed"`
	DocumentID string `json:"document_id"`
	Href       string `json:"href,omitempty,omitzero"`
	Language   string `json:"language"`
	UserID     int    `json:"user_id,omitempty,omitzero"`
}
type Type string

const (
	BankCard      Type = "bank_card"
	Blockquote    Type = "blockquote"
	Bold          Type = "bold"
	BotCommand    Type = "bot_command"
	Cashtag       Type = "cashtag"
	Code          Type = "code"
	CustomEmoji   Type = "custom_emoji"
	Email         Type = "email"
	Hashtag       Type = "hashtag"
	Italic        Type = "italic"
	Link          Type = "link"
	Mention       Type = "mention"
	MentionName   Type = "mention_name"
	Phone         Type = "phone"
	Plain         Type = "plain"
	Pre           Type = "pre"
	Spoiler       Type = "spoiler"
	Strikethrough Type = "strikethrough"
	TextLink      Type = "text_link"
	Underline     Type = "underline"
	Unknown       Type = "unknown"
)

func (e TextEntity) String() string {
	switch e.Type {
	// case BankCard:
	case Blockquote:
		return fmt.Sprintf("Citation(%s)", e.Text)
	case Bold:
		return fmt.Sprintf("Bold(%s)", e.Text)
	case Italic:
		return fmt.Sprintf("Italic(%s)", e.Text)
	case MentionName:
		return fmt.Sprintf("%s userid(%d)", e.Text, e.UserID)
	case
		BankCard,
		BotCommand,
		Cashtag, // Text prefixed with a dollar sign, such as $PATH.
		Code,
		CustomEmoji,
		Email,
		Hashtag,
		Link,
		Mention,
		Phone,
		Plain,
		Unknown: // Telegram interprets this as a date and offers to create a reminder.
		return e.Text
	case Pre: // Code block.
		return e.Language + ": " + e.Text
	case Spoiler:
		return fmt.Sprintf("Spoiler(%s)", e.Text)
	case Strikethrough:
		return fmt.Sprintf("Strikethrough(%s)", e.Text)
	case TextLink:
		return fmt.Sprintf("%s: %s", e.Text, e.Href)
	case Underline:
		return fmt.Sprintf("Underline(%s)", e.Text)
	default:
		slog.Warn("Type is undefined. Fix required!")

		b, err := json.Marshal(e)
		if err != nil {
			// ignore
		}

		return string(b)
	}
}
