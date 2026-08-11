package text

import (
	"encoding/json"
	"log/slog"
	"strings"

	"telegram-export-parser/domain/part"
)

type Text struct {
	Parts []part.Part
}

func (t *Text) String() string {
	parts := make([]string, 0, len(t.Parts))

	for _, part := range t.Parts {
		parts = append(parts, part.String())
	}

	return strings.Join(parts, "")
}

func (t *Text) UnmarshalJSON(data []byte) error {
	var text any

	err := json.Unmarshal(data, &text)
	if err != nil {
		return err
	}

	var parts []part.Part

	switch textValue := text.(type) {
	case string:
		part := part.Part{Type: part.TextType, M: map[part.PartType]any{part.TextType: textValue}}

		parts = append(parts, part)
	case []any:
		textParts := textValue

		for _, textPart := range textParts {
			switch vvv := textPart.(type) {
			case string:
				part := part.Part{Type: part.TextType, M: map[part.PartType]any{part.TextType: vvv}}

				parts = append(parts, part)
			case map[string]any:
				_type, ok := vvv["type"].(string)
				if !ok {
					slog.Error("failed to parse type from text")
				}

				m := make(map[part.PartType]any)

				for k1, v1 := range vvv {
					t := part.GetPartType(k1)
					if t == part.TypeType {
						continue
					}

					m[t] = v1
				}

				part := part.Part{
					Type: part.PartType(_type),
					M:    m,
				}

				parts = append(parts, part)
			default:
				slog.Error("Unknown part type in text")
			}
		}
	default:
		slog.Error("Unknown part type in text")
	}

	t.Parts = parts

	return nil
}
