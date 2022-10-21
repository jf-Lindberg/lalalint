/*
Copyright © 2022 Filip Lindberg fili21@student.bth.se
*/

package linter

import (
	"fmt"
	"regexp"
)

type BadCommentError struct {
	Row   int
	Index int
	Fault string
}

func (e BadCommentError) Error() string {
	return fmt.Sprintf("no space after comment on row %d, character %d: '%s'", e.Row, e.Index+1, e.Fault)
}

func commentEscaped(index int, lineContent string) bool {
	if index != 0 {
		return lineContent[index-1:index] == "\\"
	}
	return false
}

func invalidCommentRegex() *regexp.Regexp {
	r := fmt.Sprintf("(?P<symbol>%s)(?P<character>%s)", "%", `[^\s"']`)
	return regexp.MustCompile(r)
}

func SpaceAfterComment(line Line) (Line, error) {
	pattern := invalidCommentRegex()
	if index := GetIndex(pattern, line); index != nil {
		if commentEscaped(index[0], line.Content) {
			return line, nil
		}
		template := "${symbol} ${character}"
		replaced := pattern.ReplaceAllString(line.Content, template)
		fault := GetFault(index[0], line)
		return Line{line.Row, replaced}, BadCommentError{line.Row, index[0], fault}
	}

	return line, nil
}
