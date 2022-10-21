package linter

import (
	"fmt"
	"regexp"
)

type NoNewLine struct {
	Row int
}

type NewlineAfterSentenceErr struct {
	Row int
}

func (e NewlineAfterSentenceErr) Error() string {
	return fmt.Sprintf("no newline after sentence on row %d", e.Row)
}

// NewlineAfterSentence uses regex to find sentences and add newlines after them.
// A sentence ending is defined as a non-whitespace character followed by a dot, exclamation or question mark.
// In order for a sentence to be detected, there needs to be whitespace after the end of sentence directly followed by a character.
// In other words, a badly formatted sequence of sentences will not be formatted. For example "Foo!bar" would be interpreted as one word.
// If the original line content is a comment, the newlines that get added will also be comments.
func NewlineAfterSentence(line Line) (Line, error) {
	content := line.Content
	pattern := regexp.MustCompile(`(?P<pre>[^\s])(?P<end>[.!?])\s+(?P<post>[\w])`)
	// Regexp object returns empty string if pattern not found
	if pattern.FindString(content) == "" {
		return line, nil
	}
	template := "${pre}${end}\n%${post}"
	cIndex := GetIndex(CommentRegex(), line)
	sIndex := GetIndex(pattern, line)
	// if no comment or comment is after sentence, change template to not include "%"
	if cIndex == nil || cIndex[0] > sIndex[0] {
		template = "${pre}${end}\n${post}"
	}
	content = pattern.ReplaceAllString(content, template)

	return Line{line.Row, content}, NewlineAfterSentenceErr{line.Row}
}
