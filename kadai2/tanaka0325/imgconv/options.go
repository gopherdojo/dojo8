package imgconv

import (
	"fmt"
	"strings"
)

// Options is type for command line options.
type Options struct {
	From   *string
	To     *string
	DryRun *bool
}

func (opt Options) Validate(allowList []string) error {
	to := strings.ToLower(*opt.To)
	from := strings.ToLower(*opt.From)
	targetExts := []string{to, from}

	for _, e := range targetExts {
		if !isInclude(allowList, e) {
			return fmt.Errorf("%s is not allowed. ext is only allowed in %s", e, allowList)
		}
	}

	return nil
}

func isInclude(list []string, w string) bool {
	for _, e := range list {
		if e == w {
			return true
		}
	}

	return false
}
