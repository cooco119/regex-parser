package matcher

type Matcher struct{}

func (m *Matcher) Match(text string, pattern string) string {
	result := ""

	textLen := len(text)
	patternLen := len(pattern)
	i := 0

	for ; i < textLen && i < patternLen; i++ {
		textChar := string(text[i])
		patternChar := string(pattern[i])
		if textChar == patternChar {
			result += textChar
		} else if patternChar == "." {
			result += textChar
		}
	}

	if i != patternLen {
		return ""
	}

	return result
}

func GetMatcher() *Matcher {
	return new(Matcher)
}
