package domain

type HistoryDTO struct {
	Name     string       `json:"name"`
	Type     string       `json:"type"`
	ID       int          `json:"id"`
	Messages []MessageDTO `json:"messages"`
}

type MessageDTO struct {
	ID                int                 `json:"id"`
	Type              string              `json:"type"`
	Date              TelegramTime        `json:"date"`
	DateUnixtime      string              `json:"date_unixtime"`
	Actor             string              `json:"actor,omitempty"`
	ActorID           string              `json:"actor_id,omitempty"`
	Action            string              `json:"action,omitempty"`
	Title             string              `json:"title,omitempty"`
	Text              Text                `json:"text"`
	TextEntities      []any               `json:"text_entities"`
	From              string              `json:"from,omitempty"`
	FromID            string              `json:"from_id,omitempty"`
	ReplyToMessageID  int                 `json:"reply_to_message_id,omitempty"`
	Photo             string              `json:"photo,omitempty"`
	PhotoFileSize     int                 `json:"photo_file_size,omitempty"`
	Width             int                 `json:"width,omitempty"`
	Height            int                 `json:"height,omitempty"`
	ForwardedFrom     string              `json:"forwarded_from,omitempty"`
	MessageID         int                 `json:"message_id,omitempty"`
	Members           []string            `json:"members,omitempty"`
	Edited            TelegramTime        `json:"edited,omitempty"`
	EditedUnixtime    string              `json:"edited_unixtime,omitempty"`
	ViaBot            string              `json:"via_bot,omitempty"`
	File              string              `json:"file,omitempty"`
	FileName          string              `json:"file_name,omitempty"`
	FileSize          int                 `json:"file_size,omitempty"`
	Thumbnail         string              `json:"thumbnail,omitempty"`
	ThumbnailFileSize int                 `json:"thumbnail_file_size,omitempty"`
	MediaType         string              `json:"media_type,omitempty"`
	StickerEmoji      string              `json:"sticker_emoji,omitempty"`
	MimeType          string              `json:"mime_type,omitempty"`
	DurationSeconds   int                 `json:"duration_seconds,omitempty"`
	InlineBotButtons  [][]InlineBotButton `json:"inline_bot_buttons,omitempty"`
	Poll              Poll                `json:"poll,omitempty"`
	Reactions         []Reaction          `json:"reactions,omitempty"`
}

type InlineBotButton struct {
	Type       string `json:"type"`
	Text       string `json:"text"`
	DataBase64 string `json:"dataBase64"`
	Data       string `json:"data"`
}

type Poll struct {
	Question    string   `json:"question"`
	Closed      bool     `json:"closed"`
	TotalVoters int      `json:"total_voters"`
	Answers     []Answer `json:"answers"`
}

type Answer struct {
	Text   string `json:"text"`
	Voters int    `json:"voters"`
	Chosen bool   `json:"chosen"`
}

type Reaction struct {
	Type    string   `json:"type"`
	Count   int      `json:"count"`
	Emoji   string   `json:"emoji"`
	Recents []Recent `json:"recent"`
}

type Recent struct {
	From   string `json:"from"`
	FromID string `json:"from_id"`
	Date   string `json:"date"`
}
