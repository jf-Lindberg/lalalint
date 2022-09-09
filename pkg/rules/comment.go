/*
Copyright © 2022 Filip Lindberg jakob.filip.lindberg@gmail.com
*/

package rules

import (
	"strings"
)

// IDEA: make this into a factory. Create "LintCharacter" with options to add a character before or after the specified char

/*
LintComment Checks if comment is formatted correctly, meaning that there should be a space after valid comments.
Valid comments meaning non-escaped ones such as "\% this is an escaped comment".
*/
func LintComment(line string) string {
	// Checks that line includes %
	if index := commentIndex(line, "%"); index != -1 {
		// First checks if index is 0, then checks that previous character is not the escape character.
		// This check is in place to avoid index out of range errors.
		if index == 0 || !commentEscaped(index, line) {
			// If comment is correct, leave and return whole line
			if !commentIsCorrect(index, line) {
				// Else, return corrected line
				return correctLine(index, line)
			}
		}
	}

	// Return unchanged line
	return line
}

/*
commentIndex returns the index of substring
*/
func commentIndex(line string, char string) int {
	return strings.Index(line, char)
}

/*
commentEscaped determines whether the LaTeX escape character was used before the comment symbol
*/
func commentEscaped(index int, line string) bool {
	return line[index-1:index] == "\\"
}

/*
commentIsCorrect determines whether there is a space after the comment symbol
*/
func commentIsCorrect(index int, line string) bool {
	return line[index+1:index+2] == " "
}

/*
correctLine corrects the line to the proper format (adds space after the %)
*/
func correctLine(index int, line string) string {
	return line[:index+1] + " " + line[index+1:]
}
