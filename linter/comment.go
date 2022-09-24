/*
Copyright © 2022 Filip Lindberg fili21@student.bth.se
*/

package linter

import (
	"fmt"
	"strings"
)

type BadCommentError struct {
	Row   int
	Index int
	Fault string
}

func (e BadCommentError) Error() string {
	return fmt.Sprintf("no space after comment on row %d, character %d: '%s'", e.Row, e.Index+1, e.Fault)
}

/*
LintComment Checks if comment is formatted correctly, meaning that there should be a space after valid comments.
Valid comments meaning non-escaped ones such as "\% this is an escaped comment".
*/
func LintComment(line Line, symbol string) (Line, error) {
	// Checks that line includes %
	if index := commentIndex(line.Content, symbol); index != -1 {
		// First checks if index is 0, then checks that previous character is not the escape character.
		// This check is in place to avoid index out of range errors and to not format comments that are escaped.
		if index == 0 || !commentEscaped(index, line.Content) {
			// If comment is correct, leave and return whole line
			if !commentIsCorrect(index, line.Content) {
				// Else, return corrected line
				fault := getFault(index, line)
				return correctLine(index, line), BadCommentError{line.Row, index, fault}
			}
		}
	}

	// Return unchanged line
	return line, nil
}

/*
commentIndex returns the index of substring
*/
func commentIndex(lineContent string, char string) int {
	return strings.Index(lineContent, char)
}

/*
commentEscaped determines whether the LaTeX escape character was used before the comment symbol
*/
func commentEscaped(index int, lineContent string) bool {
	return lineContent[index-1:index] == "\\"
}

/*
commentIsCorrect determines whether there is a space after the comment symbol
*/
func commentIsCorrect(index int, lineContent string) bool {
	return lineContent[index+1:index+2] == " "
}

/*
correctLine corrects the line to the proper format (adds space after the %)
*/
func correctLine(index int, line Line) Line {
	content := line.Content[:index+1] + " " + line.Content[index+1:]
	return Line{line.Row, content}
}

/*
getFault gets the context of the error and returns it
*/
func getFault(index int, line Line) string {
	var fault string
	end := len(line.Content)
	if end-index > 5 {
		fault = line.Content[index : index+5]
	} else {
		fault = line.Content[index : end-1]
	}
	return fault
}
