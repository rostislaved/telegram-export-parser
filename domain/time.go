package domain

import (
	"encoding/json"
	"errors"
	"time"
)

type TelegramTime struct {
	time.Time
}

func (t *TelegramTime) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		t.Time = time.Time{}

		return nil
	}

	var s string

	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	if s == "" {
		t.Time = time.Time{}

		return nil
	}

	tt, err := time.Parse(time.RFC3339Nano, s)
	if err == nil {
		t.Time = tt

		return nil
	}

	tt, err = time.Parse(time.RFC3339, s)
	if err == nil {
		t.Time = tt

		return nil
	}

	tt, err = time.ParseInLocation("2006-01-02T15:04:05", s, time.UTC)
	if err == nil {
		t.Time = tt

		return nil
	}

	return errors.New("unsupported time format: " + s)
}

func (t TelegramTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}

	return json.Marshal(t.Format(time.RFC3339Nano))
}
