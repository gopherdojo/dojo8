package imgconv

// Args is type for command line arguments.
type Args []string

func (args Args) uniq() []string {
	m := map[string]bool{}
	u := []string{}

	for _, v := range args {
		if !m[v] {
			m[v] = true
			u = append(u, v)
		}
	}

	return u
}
