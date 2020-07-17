package imgconv

import (
	"errors"
	"fmt"
	"strings"
)

// Args is type for command line options.
type Options struct {
	From   *string
	To     *string
	DryRun *bool
}

func (opt Options) validate(allow_list []string) error {
	to := strings.ToLower(*opt.To)
	from := strings.ToLower(*opt.From)
	targetExts := []string{to, from}
	for _, e := range targetExts {
		if err := include(allow_list, e); err != nil {
			return fmt.Errorf("%w. ext is only allowd in %s", err, allow_list)
		}
	}

	return nil
}

func include(list []string, w string) error {
	for _, e := range list {
		if e == w {
			return nil
		}
	}
	return errors.New(w + " is not allowed")
}
