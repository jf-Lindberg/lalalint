package linter

import (
	"fmt"
	"regexp"
	"strings"
)

// NoBlankLineErr is the error returned if a linter problem is found.
type NoBlankLineErr struct {
	Row int
}

// Error message.
func (e NoBlankLineErr) Error() string {
	return fmt.Sprintf("no blank line(s) before section on row %d", e.Row)
}

// addBlankLines adds blank lines to the input line content and returns a Line struct with the new content.
// If the amount of blank lines already is at the configured amount or higher prior to the input Line, this returns the original Line struct.
func addBlankLines(blankLinesCfg int, blankLinesBefore int, line Line) (Line, error) {
	remainder := blankLinesCfg - blankLinesBefore
	if remainder <= 0 {
		return line, nil
	}
	content := strings.Repeat("\n", remainder) + line.Content
	err := NoBlankLineErr{line.Row}
	return Line{line.Row, content}, err
}

// BlankLinesBeforeSection checks if the input Line contains a non-commented section and calls addBlankLines if it does.
// It takes three arguments:
// The input Line.
// blankLinesCfg which is the amount of blank lines that should be added before a section according to config.
// blankLinesBefore which is the amount of blank lines that are present before the input Line. This is tracked in the lint function.
func BlankLinesBeforeSection(blankLinesCfg int, blankLinesBefore int, line Line) (Line, error) {
	pattern := regexp.MustCompile(`^\\(?P<section>(sub)?section|chapter)`)
	if isSection(line.Content, pattern) {
		cIndex := GetIndex(CommentRegex(), line)
		sIndex := GetIndex(pattern, line)
		if cIndex == nil || cIndex[0] > sIndex[0] {
			return addBlankLines(blankLinesCfg, blankLinesBefore, line)
		}
	}

	return line, nil
}

func isSection(content string, r *regexp.Regexp) bool {
	return r.Match([]byte(content))
}
