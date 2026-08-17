package text

import (
	"encoding/json"
	"strings"

	"telegram-export-parser/internal/app/domain/text_entity"
)

type Text struct {
	TextEntities []text_entity.TextEntity
}

func (t *Text) String() string {
	parts := make([]string, 0, len(t.TextEntities))

	for _, textPart := range t.TextEntities {
		parts = append(parts, textPart.Text)
	}

	return strings.Join(parts, "")
}

// Text can be either a string or an array of elements.
// An array element can be either a string or a text_entity.TextEntity.

func (t *Text) UnmarshalJSON(data []byte) error {
	var plainText string

	err := json.Unmarshal(data, &plainText)
	if err == nil {
		textPart := text_entity.TextEntity{
			Type: text_entity.Plain,
			Text: plainText,
		}

		t.TextEntities = []text_entity.TextEntity{textPart}

		return nil
	}

	var rawParts []json.RawMessage

	err = json.Unmarshal(data, &rawParts)
	if err != nil {
		return err
	}

	parts := make([]text_entity.TextEntity, 0, len(rawParts))

	for _, rawPart := range rawParts {
		err = json.Unmarshal(rawPart, &plainText)
		if err == nil {
			textPart := text_entity.TextEntity{
				Type: text_entity.Plain,
				Text: plainText,
			}

			parts = append(parts, textPart)

			continue
		}

		var entity text_entity.TextEntity

		err = json.Unmarshal(rawPart, &entity)
		if err != nil {
			return err
		}

		parts = append(parts, entity)
	}

	t.TextEntities = parts

	return nil
}
