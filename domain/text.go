package domain

import (
	"encoding/json"
	"log/slog"
	"strings"
)

type Text struct {
	Parts []Part
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

	var parts []Part

	switch textValue := text.(type) {
	case string:
		part := Part{Type: TextType, M: map[PartType]any{TextType: textValue}}

		parts = append(parts, part)
	case []any:
		textParts := textValue

		for _, textPart := range textParts {
			switch vvv := textPart.(type) {
			case string:
				part := Part{Type: TextType, M: map[PartType]any{TextType: vvv}}

				parts = append(parts, part)
			case map[string]any:
				_type, ok := vvv["type"].(string)
				if !ok {
					slog.Error("failed to parse type from text")
				}

				m := make(map[PartType]any)

				for k1, v1 := range vvv {
					t := getPartType(k1)
					if t == TypeType {
						continue
					}

					m[t] = v1
				}

				part := Part{
					Type: PartType(_type),
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
