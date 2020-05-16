package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"regex-parser/matcher"
)

var mat = new(matcher.Matcher)

func TestMatcherShouldMatchSameString(t *testing.T) {
	text := "abcd"
	pattern := "abcd"

	result := mat.Match(text, pattern)

	assert.Equal(t, text, result, "Pattern same as text should return text itself")
}

func TestMatcherShouldMatchSingleDot(t *testing.T) {
	text := "abcd"

	pattern := ".bcd"
	result := mat.Match(text, pattern)
	assert.Equal(t, text, result)

	pattern = "a.cd"
	result = mat.Match(text, pattern)
	assert.Equal(t, text, result)

	pattern = "ab.d"
	result = mat.Match(text, pattern)
	assert.Equal(t, text, result)

	pattern = "abc."
	result = mat.Match(text, pattern)
	assert.Equal(t, text, result)
}

func TestMatcherShouldFailForDiffrentLength(t *testing.T) {
	text := "abcd"
	pattern := "abcde"
	result := mat.Match(text, pattern)

	assert.Equal(t, "", result)
}
