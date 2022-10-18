package linter

import (
	"fmt"
	"github.com/spf13/viper"
	"regexp"
	"strings"
)

type NoNewLine struct {
	Row int
}

func (e NoNewLine) Error() string {
	return fmt.Sprintf("no newline after sentence on row %d", e.Row)
}

type NewLine struct {
	line Line
	err  error
}

func NewlineAfterSentence(line Line) []NewLine {
	arr := make([]NewLine, 0)
	content := line.Content
	pattern := regexp.MustCompile("(?P<pre>[^\\s])(?P<end>[.!?])\\s+(?P<post>[A-Z])")
	// Regexp object returns empty string if pattern not found
	if pattern.FindString(content) == "" || !viper.GetBool("rules.newlineaftersentence.enabled") {
		arr = append(arr, NewLine{line, nil})
		return arr
	}

	splitLines := splitLines(pattern, line)
	row := line.Row
	for i, _ := range splitLines {
		newline := Line{row, splitLines[i]}
		arr = append(arr, NewLine{newline, NoNewLine{row}})
		row++
	}

	return arr
}

func splitLines(pattern *regexp.Regexp, line Line) []string {
	template := "${pre}${end}\n%${post}"
	cIndex := GetIndex(CommentRegex(), line)
	sIndex := GetIndex(pattern, line)

	if cIndex == nil || cIndex[0] > sIndex[0] {
		template = "${pre}${end}\n${post}"
	}
	replaced := pattern.ReplaceAllString(line.Content, template)

	return strings.Split(replaced, "\n")
}
