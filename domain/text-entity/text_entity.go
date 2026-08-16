package text_entity

type TextEntity struct {
	Type       TextEntityType `json:"type"`
	Text       string         `json:"text"`
	Collapsed  bool           `json:"collapsed"`
	DocumentID string         `json:"document_id"`
	Href       string         `json:"href,omitempty,omitzero"`
	Language   string         `json:"language"`
	UserID     int            `json:"user_id,omitempty,omitzero"`
}
type TextEntityType string

const (
	Blockquote    TextEntityType = "blockquote"
	Bold          TextEntityType = "bold"
	BotCommand    TextEntityType = "bot_command"
	Cashtag       TextEntityType = "cashtag"
	Code          TextEntityType = "code"
	CustomEmoji   TextEntityType = "custom_emoji"
	Email         TextEntityType = "email"
	Hashtag       TextEntityType = "hashtag"
	Italic        TextEntityType = "italic"
	Link          TextEntityType = "link"
	Mention       TextEntityType = "mention"
	MentionName   TextEntityType = "mention_name"
	Phone         TextEntityType = "phone"
	Plain         TextEntityType = "plain"
	Pre           TextEntityType = "pre"
	Spoiler       TextEntityType = "spoiler"
	Strikethrough TextEntityType = "strikethrough"
	TextLink      TextEntityType = "text_link"
	Underline     TextEntityType = "underline"
)
