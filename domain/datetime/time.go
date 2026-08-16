package datetime

import (
	"encoding/json"
	"errors"
	"time"
)

// Telegram stores datetime as "2024-07-28T16:32:57"
type DateTime struct {
	time.Time
}

const telegramDateTimeLayout = "2006-01-02T15:04:05"

func (t *DateTime) UnmarshalJSON(b []byte) error {
	s := string(b[1 : len(b)-1]) // trim quotes

	tt, err := time.ParseInLocation(telegramDateTimeLayout, s, time.UTC)
	if err != nil {
		return errors.New("unsupported time format: " + s)
	}

	t.Time = tt

	return nil
}

func (t DateTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}

	return json.Marshal(t.Format(time.RFC3339Nano))
}
