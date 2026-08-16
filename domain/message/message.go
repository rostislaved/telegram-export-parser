package message

import (
	"telegram-export-parser/domain/datetime"
	"telegram-export-parser/domain/text"
	text_entity "telegram-export-parser/domain/text-entity"
)

type Message struct {
	Action                    string                   `json:"action,omitempty,omitzero"`
	Actor                     *string                  `json:"actor,omitempty,omitzero"`
	ActorID                   string                   `json:"actor_id,omitempty,omitzero"`
	Address                   string                   `json:"address,omitempty,omitzero"`
	Boosts                    int                      `json:"boosts,omitempty,omitzero"`
	ContactInformation        *ContactInformation      `json:"contact_information,omitempty,omitzero"`
	ContactVcard              string                   `json:"contact_vcard,omitempty,omitzero"`
	ContactVcardFileSize      int                      `json:"contact_vcard_file_size,omitempty,omitzero"`
	Date                      datetime.DateTime        `json:"date"`
	DateUnixtime              string                   `json:"date_unixtime"`
	DurationSeconds           int                      `json:"duration_seconds,omitzero"`
	Edited                    datetime.DateTime        `json:"edited,omitzero"`
	EditedUnixtime            string                   `json:"edited_unixtime,omitzero"`
	File                      string                   `json:"file,omitzero"`
	FileName                  string                   `json:"file_name,omitzero"`
	FileSize                  int                      `json:"file_size,omitzero"`
	ForwardedFrom             *string                  `json:"forwarded_from,omitzero"`
	ForwardedFromID           string                   `json:"forwarded_from_id,omitempty,omitzero"`
	From                      *string                  `json:"from,omitzero"`
	FromID                    string                   `json:"from_id,omitzero"`
	GameDescription           string                   `json:"game_description,omitempty,omitzero"`
	GameLink                  string                   `json:"game_link,omitempty,omitzero"`
	GameMessageID             int                      `json:"game_message_id,omitempty,omitzero"`
	GameTitle                 string                   `json:"game_title,omitempty,omitzero"`
	GiveawayInformation       *GiveawayInformation     `json:"giveaway_information,omitempty,omitzero"`
	Height                    int                      `json:"height,omitzero"`
	ID                        int                      `json:"id"`
	InlineBotButtons          [][]InlineBotButton      `json:"inline_bot_buttons,omitzero"`
	Inviter                   string                   `json:"inviter,omitempty,omitzero"`
	InvoiceInformation        string                   `json:"invoice_information,omitempty,omitzero"`
	LiveLocationPeriodSeconds int                      `json:"live_location_period_seconds,omitempty,omitzero"`
	LocationInformation       *LocationInformation     `json:"location_information,omitempty,omitzero"`
	MediaSpoiler              bool                     `json:"media_spoiler,omitempty,omitzero"`
	MediaType                 string                   `json:"media_type,omitzero"`
	Members                   []*string                `json:"members,omitempty"`
	MessageID                 int                      `json:"message_id,omitempty,omitzero"`
	MimeType                  string                   `json:"mime_type,omitzero"`
	NewIconEmojiID            int                      `json:"new_icon_emoji_id,omitempty,omitzero"`
	NewTitle                  string                   `json:"new_title,omitempty,omitzero"`
	Performer                 string                   `json:"performer,omitempty,omitzero"`
	Photo                     string                   `json:"photo,omitzero"`
	PhotoFileSize             int                      `json:"photo_file_size,omitzero"`
	PlaceName                 string                   `json:"place_name,omitempty,omitzero"`
	Poll                      *Poll                    `json:"poll,omitzero"`
	Reactions                 []Reaction               `json:"reactions,omitzero"`
	ReplyToMessageID          int                      `json:"reply_to_message_id,omitzero"`
	ReplyToPeerID             string                   `json:"reply_to_peer_id,omitempty,omitzero"`
	RichMessage               *RichMessage             `json:"rich_message,omitempty,omitzero"`
	SavedFrom                 string                   `json:"saved_from,omitempty,omitzero"`
	Score                     int                      `json:"score,omitempty,omitzero"`
	StickerEmoji              string                   `json:"sticker_emoji,omitzero"`
	Text                      text.Text                `json:"text"`
	TextEntities              []text_entity.TextEntity `json:"text_entities"`
	Thumbnail                 string                   `json:"thumbnail,omitzero"`
	ThumbnailFileSize         int                      `json:"thumbnail_file_size,omitzero"`
	Title                     string                   `json:"title,omitzero"`
	Type                      Type                     `json:"type"` // service type - pin_message, "user1 Added user2", etc
	ViaBot                    string                   `json:"via_bot,omitzero"`
	Width                     int                      `json:"width,omitzero"`
}

type Type string

const (
	Blockquote        Type = "blockquote"
	Bold                   = "bold"
	BotCommand             = "bot_command"
	Code                   = "code"
	CustomEmoji            = "custom_emoji"
	Email                  = "email"
	Emoji                  = "emoji"
	Hashtag                = "hashtag"
	Italic                 = "italic"
	Link                   = "link"
	Mention                = "mention"
	MentionName            = "mention_name"
	MessageT               = "message"
	Phone                  = "phone"
	Plain                  = "plain"
	Pre                    = "pre"
	PrivateSupergroup      = "private_supergroup"
	Service                = "service"
	Spoiler                = "spoiler"
	Strikethrough          = "strikethrough"
	TextLink               = "text_link"
	Underline              = "underline"
)

type ContactInformation struct {
	FirstName   string `json:"first_name,omitzero"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number,omitzero"`
}

type GiveawayInformation struct {
	AdditionalPrize      string `json:"additional_prize"`
	Channels             []int  `json:"channels"`
	Countries            []any  `json:"countries"`
	IsOnlyNewSubscribers bool   `json:"is_only_new_subscribers"`
	Months               int    `json:"months,omitzero"`
	Quantity             int    `json:"quantity,omitzero"`
	Stars                int    `json:"stars"`
	UntilDate            string `json:"until_date,omitzero"`
}
type LocationInformation struct {
	Latitude  float64 `json:"latitude,omitzero"`
	Longitude float64 `json:"longitude,omitzero"`
}

type InlineBotButton struct {
	ButtonID   int    `json:"button_id,omitempty,omitzero"`
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
	Text   text.Text `json:"text"`
	Voters int       `json:"voters"`
	Chosen bool      `json:"chosen"`
}

type Reaction struct {
	Type       string   `json:"type"` // "emoji" or "custom_emoji"
	DocumentID string   `json:"document_id"`
	Count      int      `json:"count"`
	Emoji      string   `json:"emoji"`
	Recent     []Recent `json:"recent"`
}

type Recent struct {
	From   *string           `json:"from"`    // Имя человека поставившего реакцию
	FromID string            `json:"from_id"` // ID человека поставившего реакцию вида: "user%ID", например user447282515
	Date   datetime.DateTime `json:"date"`
}

type RichMessage struct {
	Blocks []RichMessageBlock `json:"blocks"`
	Part   bool               `json:"part"`
	Rtl    bool               `json:"rtl"`
}

type RichMessageBlock struct {
	Blocks    []RichMessageNestedBlock `json:"blocks,omitempty"`
	Caption   *RichMessageCaption      `json:"caption,omitempty,omitzero"`
	Content   string                   `json:"content,omitempty,omitzero"`
	Items     []RichMessageItem        `json:"items,omitempty"`
	Kind      string                   `json:"kind,omitempty,omitzero"`
	Pullquote bool                     `json:"pullquote"`
	Text      *RichMessageText         `json:"text,omitempty,omitzero"`
	Type      string                   `json:"type,omitzero"`
}

type RichMessageNestedBlock struct {
	Caption   *RichMessageCaption `json:"caption,omitempty,omitzero"`
	Content   string              `json:"content,omitempty,omitzero"`
	Language  string              `json:"language,omitempty,omitzero"`
	Pullquote bool                `json:"pullquote"`
	Text      RichMessageAnyText  `json:"text"`
	Type      string              `json:"type,omitzero"`
}

type RichMessageCaption struct {
	Type string `json:"type,omitzero"`
}

type RichMessageAnyText struct {
	Text any    `json:"text"`
	Type string `json:"type,omitzero"`
}

type RichMessageText struct {
	Text any    `json:"text,omitempty"`
	Type string `json:"type,omitzero"`
}

type RichMessageItem struct {
	Content   string              `json:"content,omitzero"`
	TaskState string              `json:"task_state,omitzero"`
	Text      RichMessageItemText `json:"text"`
}

type RichMessageItemText struct {
	Text string `json:"text,omitzero"`
	Type string `json:"type,omitzero"`
}
