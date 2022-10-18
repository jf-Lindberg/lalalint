/*
Copyright © 2022 Filip Lindberg jakob.filip.lindberg@gmail.com
*/

package linter

import (
	"testing"
)

type indentTest struct {
	indentLevel int
	tabs        int
	input       Line
	expected    Line
}

// IndentTests are the test cases called by TestIndentEnvironments
var IndentTests = []indentTest{
	indentTest{indentLevel: 0, tabs: 1, input: Line{0, "Hello!"}, expected: Line{0, "Hello!"}},
	indentTest{indentLevel: 1, tabs: 1, input: Line{0, "Hello!"}, expected: Line{0, "\tHello!"}},
	indentTest{indentLevel: 0, tabs: 1, input: Line{0, "\\begin"}, expected: Line{0, "\\begin"}},
	indentTest{indentLevel: 2, tabs: 1, input: Line{0, "\\begin"}, expected: Line{0, "\t\t\\begin"}},
	indentTest{indentLevel: 0, tabs: 1, input: Line{0, "\\end"}, expected: Line{0, "\\end"}},
	//indentTest{indentLevel: 3, tabs: 1, input: Line{0, "\\end"}, expected: Line{0, "\t\t\\end"}},
}

func TestIndentEnvironments(t *testing.T) {
	for _, test := range IndentTests {
		if output := IndentEnvironments(test.indentLevel, test.tabs, test.input); output != test.expected {
			t.Errorf("Output %q not equal to expected %q. Types: %T %T", output, test.expected, output, test.expected)
		}
	}
}
