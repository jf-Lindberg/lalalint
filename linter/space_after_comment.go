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

///*
//SpaceAfterComment Checks if comment is formatted correctly, meaning that there should be a space after valid comments.
//Valid comments meaning non-escaped ones such as "\% this is an escaped comment".
//*/
//func SpaceAfterComment(line Line, symbol string) (Line, error) {
//	// Regex: .*%[^\s] or maybe even better (%)([^\s]), then just add whitespace between CG1 and CG2
//	// Checks that line includes %
//	if index := commentIndex(line.Content, symbol); index != -1 {
//		// First checks if index is 0, then checks that previous character is not the escape character.
//		// This check is in place to avoid index out of range errors and to not format comments that are escaped.
//		if index == 0 || !commentEscaped(index, line.Content) {
//			// If comment is correct, leave and return whole line
//			if !commentIsCorrect(index, line.Content) {
//				// Else, return corrected line
//				fault := GetFault(index, line)
//				return correctLine(index, line), BadCommentError{line.Row, index, fault}
//			}
//		}
//	}
//
//	// Return unchanged line
//	return line, nil
//}

/*
commentEscaped determines whether the LaTeX escape character was used before the comment symbol
*/
func commentEscaped(index int, lineContent string) bool {
	if index != 0 {
		return lineContent[index-1:index] == "\\"
	}
	return false
}

//func commentIndex(line Line) []int {
//	regex := invalidCommentRegex()
//	return regex.FindIndex([]byte(line.Content))
//}

func invalidCommentRegex() *regexp.Regexp {
	r := fmt.Sprintf("(?P<symbol>%s)(?P<character>%s)", "%", `[^\s]`)
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

//func FindComment(symbol string, content string) int {
//	r := fmt.Sprintf("(?P<symbol>%s(?P<character>%s)", symbol, `[^\s]`)
//	pattern := regexp.MustCompile(r)
//	match := pattern.FindIndex([]byte(content))
//	return
//}

///*
//commentIndex returns the index of substring
//*/
//func commentIndex(lineContent string, char string) int {
//	return strings.Index(lineContent, char)
//}

///*
//commentIsCorrect determines whether there is a space after the comment symbol
//*/
//func commentIsCorrect(index int, lineContent string) bool {
//	return lineContent[index+1:index+2] == " "
//}

///*
//correctLine corrects the line to the proper format (adds space after the %)
//*/
//func correctLine(index int, line Line) Line {
//	content := line.Content[:index+1] + " " + line.Content[index+1:]
//	return Line{line.Row, content}
//}
