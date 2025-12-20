package text_entity

type TextEntity string

const (
	Blockquote    TextEntity = "blockquote"
	Bold          TextEntity = "bold"
	BotCommand    TextEntity = "bot_command"
	Cashtag       TextEntity = "cashtag"
	Code          TextEntity = "code"
	CustomEmoji   TextEntity = "custom_emoji"
	Email         TextEntity = "email"
	Hashtag       TextEntity = "hashtag"
	Italic        TextEntity = "italic"
	Link          TextEntity = "link"
	Mention       TextEntity = "mention"
	MentionName   TextEntity = "mention_name"
	Phone         TextEntity = "phone"
	Plain         TextEntity = "plain"
	Pre           TextEntity = "pre"
	Spoiler       TextEntity = "spoiler"
	Strikethrough TextEntity = "strikethrough"
	TextLink      TextEntity = "text_link"
	Underline     TextEntity = "underline"
)
