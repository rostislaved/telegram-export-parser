package domain

type Part struct {
	Type PartType
	M    map[PartType]any
}

type PartType string

const (
	Blockquote    PartType = "blockquote"
	Bold          PartType = "bold"
	BotCommand    PartType = "bot_command"
	Cashtag       PartType = "cashtag"
	Code          PartType = "code"
	CustomEmoji   PartType = "custom_emoji"
	Email         PartType = "email"
	Hashtag       PartType = "hashtag"
	Italic        PartType = "italic"
	Link          PartType = "link"
	Mention       PartType = "mention"
	MentionName   PartType = "mention_name"
	Phone         PartType = "phone"
	Pre           PartType = "pre"
	Spoiler       PartType = "spoiler"
	Strikethrough PartType = "strikethrough"
	TextLink      PartType = "text_link"
	TextType      PartType = "text" // в TextEntity это судя по всему Plain
	TypeType      PartType = "type"
	Underline     PartType = "underline"

	Language   PartType = "language"
	Href       PartType = "href"
	DocumentID PartType = "document_id"
	Collapsed  PartType = "collapsed"
	UserID     PartType = "user_id"
)

func (p Part) String() string {
	switch p.Type {
	case TextType:
		return p.M["text"].(string)
	case Mention:
		return p.M["text"].(string)
	case Hashtag:
		return p.M["text"].(string)
	case Link:
		return p.M["text"].(string)
	case BotCommand:
		return p.M["text"].(string)
	case Bold:
		return p.M["text"].(string)
	case Code:
		return p.M["text"].(string)
	case Pre:
		return p.M["text"].(string)
	case Phone:
		return p.M["text"].(string)
	case Italic:
		return p.M["text"].(string)
	case MentionName:
		return p.M["text"].(string)
	case Email:
		return p.M["text"].(string)
	case TextLink:
		return p.M["text"].(string)
	case Cashtag:
		return p.M["text"].(string)
	case Underline:
		return p.M["text"].(string)
	case Strikethrough:
		return p.M["text"].(string)
	case Spoiler:
		return p.M["text"].(string)
	case CustomEmoji:
		return p.M["text"].(string)
	case Blockquote:
		return p.M["text"].(string)
	default:
		panic("unknown part type: " + p.Type)
	}
}

func getPartType(s string) PartType {
	switch s {
	case "type":
		return TypeType
	case "text":
		return TextType
	case "mention":
		return Mention
	case "hashtag":
		return Hashtag
	case "link":
		return Link
	case "bot_command":
		return BotCommand
	case "bold":
		return Bold
	case "code":
		return Code
	case "pre":
		return Pre
	case "phone":
		return Phone
	case "italic":
		return Italic
	case "mention_name":
		return MentionName
	case "email":
		return Email
	case "text_link":
		return TextLink
	case "cashtag":
		return Cashtag
	case "underline":
		return Underline
	case "strikethrough":
		return Strikethrough
	case "spoiler":
		return Spoiler
	case "custom_emoji":
		return CustomEmoji
	case "blockquote":
		return Blockquote
	case "language":
		return Language
	case "user_id":
		return UserID
	case "href":
		return Href
	case "document_id":
		return DocumentID
	case "collapsed":
		return Collapsed
	default:
		panic("unknown part type: " + s)
	}
}
