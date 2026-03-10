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
	Actor             string              `json:"actor,omitzero"`
	ActorID           string              `json:"actor_id,omitzero"`
	Action            string              `json:"action,omitzero"`
	Title             string              `json:"title,omitzero"`
	Text              Text                `json:"text"`
	TextEntities      []any               `json:"text_entities"`
	From              string              `json:"from,omitzero"`
	FromID            string              `json:"from_id,omitzero"`
	ReplyToMessageID  int                 `json:"reply_to_message_id,omitzero"`
	Photo             string              `json:"photo,omitzero"`
	PhotoFileSize     int                 `json:"photo_file_size,omitzero"`
	Width             int                 `json:"width,omitzero"`
	Height            int                 `json:"height,omitzero"`
	ForwardedFrom     string              `json:"forwarded_from,omitzero"`
	MessageID         int                 `json:"message_id,omitzero"`
	Members           []string            `json:"members,omitzero"`
	Edited            TelegramTime        `json:"edited,omitzero"`
	EditedUnixtime    string              `json:"edited_unixtime,omitzero"`
	ViaBot            string              `json:"via_bot,omitzero"`
	File              string              `json:"file,omitzero"`
	FileName          string              `json:"file_name,omitzero"`
	FileSize          int                 `json:"file_size,omitzero"`
	Thumbnail         string              `json:"thumbnail,omitzero"`
	ThumbnailFileSize int                 `json:"thumbnail_file_size,omitzero"`
	MediaType         string              `json:"media_type,omitzero"`
	StickerEmoji      string              `json:"sticker_emoji,omitzero"`
	MimeType          string              `json:"mime_type,omitzero"`
	DurationSeconds   int                 `json:"duration_seconds,omitzero"`
	InlineBotButtons  [][]InlineBotButton `json:"inline_bot_buttons,omitzero"`
	Poll              Poll                `json:"poll,omitzero"`
	Reactions         []Reaction          `json:"reactions,omitzero"`
}

type InlineBotButton struct {
	Type       string `json:"type"`
	Text       string `json:"text"`
	DataBase64 string `json:"dataBase64"` //nolint:tagliatelle //telegram use this
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
