package piscine

func Split(s, sep string) []string {
	if sep == "" {
		return []string{s}
	}
	var slr []string
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == sep[0] && i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			if i != start && s[start:i] != "" {
				slr = append(slr, s[start:i])
			}
			start = i + len(sep)
			if start > len(s) {
				return slr
			}
			if s[i:start] == sep {
				i = start - 1
			}
		}
	}
	slr = append(slr, s[start:])
	return slr
}
