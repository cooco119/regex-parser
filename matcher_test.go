package matcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var matcher = new(Matcher)

func TestMatcherShouldMatchSameString(t *testing.T) {
	text := "abcd"
	pattern := "abcd"

	result := matcher.Match(text, pattern)

	assert.Equal(t, text, result, "Pattern same as text should return text itself")
}

func TestMatcherShouldMatchSingleDot(t *testing.T) {
	text := "abcd"

	pattern := ".bcd"
	result := matcher.Match(text, pattern)
	assert.Equal(t, text, result)

	pattern = "a.cd"
	result = matcher.Match(text, pattern)
	assert.Equal(t, text, result)

	pattern = "ab.d"
	result = matcher.Match(text, pattern)
	assert.Equal(t, text, result)

	pattern = "abc."
	result = matcher.Match(text, pattern)
	assert.Equal(t, text, result)
}

func TestMatcherShouldFailForDiffrentLength(t *testing.T) {
	text := "abcd"
	pattern := "abcde"
	result := matcher.Match(text, pattern)

	assert.Equal(t, "", result)
}
