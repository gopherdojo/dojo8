package imgconv

import (
	"testing"
)

func TestOption_validate(t *testing.T) {
	notAllowdExt := "not_allowed_ext"
	jpg := "jpg"
	png := "png"
	allowedList := []string{jpg, png}

	tests := []struct {
		name    string
		options Options
		args    []string
		isErr   bool
	}{
		{
			name: "err: Options.From is not allowed",
			options: Options{
				From: &notAllowdExt,
				To:   &png,
			},
			args:  allowedList,
			isErr: true,
		},
		{
			name: "err: Options.To is not allowed",
			options: Options{
				From: &jpg,
				To:   &notAllowdExt,
			},
			args:  allowedList,
			isErr: true,
		},
		{
			name: "ok",
			options: Options{
				From: &jpg,
				To:   &png,
			},
			args:  allowedList,
			isErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.options.validate(tt.args)
			if (tt.isErr && err == nil) || (!tt.isErr && err != nil) {
				t.Errorf("expect err is %t, but got err is %s", tt.isErr, err)
			}
		})
	}
}
