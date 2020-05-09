package matcher

type Matcher struct{}

func (m *Matcher) Match(text string, pattern string) string {
	if text == pattern {
		return text
	}

	return ""
}

func GetMatcher() *Matcher {
	return new(Matcher)
}
