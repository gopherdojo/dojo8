package cli

import (
	"testing"
)

const dummyDir = "dummy"

func TestArgs(t *testing.T) {
	tests := []struct{ name, fromFormat, toFormat, expectedMessage string }{
		{name: "valid arguments", fromFormat: "jpeg", toFormat: "png", expectedMessage: ""},
		{name: "invalid from format", fromFormat: "hoge", toFormat: "jpg", expectedMessage: "argument of \"-f, --from\" is not valid file format. invalid format: hoge"},
		{name: "invalid to format", fromFormat: "gif", toFormat: "hage", expectedMessage: "argument of \"-t, --to\" is not valid file format. invalid format: hage"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := NewArgs(dummyDir, tt.fromFormat, tt.toFormat)
			actual := args.Validate()
			if actual == nil {
				if tt.expectedMessage != "" {
					t.Errorf("want %s, got %s", actual, tt.expectedMessage)
				}
			} else if actual.Error() != tt.expectedMessage {
				t.Errorf("want %s, got %s", actual, tt.expectedMessage)
			}
		})
	}
}
