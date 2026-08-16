package message

import (
	"encoding/json"
	"testing"
)

func TestAnswerUnmarshalText(t *testing.T) {
	tests := []struct {
		name string
		data string
		want string
	}{
		{
			name: "string",
			data: `{"text":"answer","voters":1,"chosen":true}`,
			want: "answer",
		},
		{
			name: "formatted parts",
			data: `{"text":["formatted ",{"type":"bold","text":"answer"}],"voters":1,"chosen":true}`,
			want: "formatted answer",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var answer Answer

			err := json.Unmarshal([]byte(test.data), &answer)
			if err != nil {
				t.Fatal(err)
			}

			got := answer.Text.String()
			if got != test.want {
				t.Fatalf("Text.String() = %q, want %q", got, test.want)
			}
		})
	}
}
