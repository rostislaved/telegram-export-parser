package part

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
	TextType      PartType = "text" // In TextEntity this appears to map to Plain.
	TypeType      PartType = "type"
	Underline     PartType = "underline"

	BankCard PartType = "bank_card"
	Unknown  PartType = "unknown"

	Language   PartType = "language"
	Href       PartType = "href"
	DocumentID PartType = "document_id"
	Collapsed  PartType = "collapsed"
	UserID     PartType = "user_id"
)

var partTypes = map[string]PartType{
	"type":          TypeType,
	"text":          TextType,
	"mention":       Mention,
	"hashtag":       Hashtag,
	"link":          Link,
	"bot_command":   BotCommand,
	"bold":          Bold,
	"code":          Code,
	"pre":           Pre,
	"phone":         Phone,
	"italic":        Italic,
	"mention_name":  MentionName,
	"email":         Email,
	"text_link":     TextLink,
	"cashtag":       Cashtag,
	"underline":     Underline,
	"strikethrough": Strikethrough,
	"spoiler":       Spoiler,
	"custom_emoji":  CustomEmoji,
	"blockquote":    Blockquote,
	"language":      Language,
	"user_id":       UserID,
	"href":          Href,
	"document_id":   DocumentID,
	"collapsed":     Collapsed,
}

var stringPartTypes = map[PartType]struct{}{
	TextType:      {},
	Mention:       {},
	Hashtag:       {},
	Link:          {},
	BotCommand:    {},
	Bold:          {},
	Code:          {},
	Pre:           {},
	Phone:         {},
	Italic:        {},
	MentionName:   {},
	Email:         {},
	TextLink:      {},
	Cashtag:       {},
	Underline:     {},
	Strikethrough: {},
	Spoiler:       {},
	CustomEmoji:   {},
	Blockquote:    {},
	BankCard:      {},
	Unknown:       {},
}

func (p Part) String() string {
	_, ok := stringPartTypes[p.Type]
	if !ok {
		panic("unknown part type: " + p.Type)
	}

	text, ok := p.M[TextType].(string)
	if !ok {
		panic("text is not a string")
	}

	return text
}

func GetPartType(s string) PartType {
	partType, ok := partTypes[s]
	if !ok {
		panic("unknown part type: " + s)
	}

	return partType
}
