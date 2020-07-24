package cli

import (
	"testing"

	"github.com/gopherdojo/dojo8/kadai2/kynefuk/helper"
)

const dummyDir = "dummy"

func TestArgs(t *testing.T) {
	cases := []struct{ name, fromFormat, toFormat, expectedMessage string }{
		{name: "valid arguments", fromFormat: "jpeg", toFormat: "png", expectedMessage: ""},
		{name: "invalid from format", fromFormat: "hoge", toFormat: "jpg", expectedMessage: "argument of \"-f, --From\" is not valid file format. invalid format: hoge"},
		{name: "invalid to format", fromFormat: "gif", toFormat: "hage", expectedMessage: "argument of \"-t, --To\" is not valid file format. invalid format: hage"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			args := NewArgs(dummyDir, c.fromFormat, c.toFormat)
			actual := args.Validate()
			if actual == nil {
				if c.expectedMessage != "" {
					helper.ErrorHelper(t, actual, c.expectedMessage)
				}
			} else if actual.Error() != c.expectedMessage {
				helper.ErrorHelper(t, actual, c.expectedMessage)
			}
		})
	}
}
