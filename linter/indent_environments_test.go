/*
Copyright © 2022 Filip Lindberg fili21@student.bth.se
*/

package linter

import (
	"github.com/jf-Lindberg/lalalint/helper"
	"testing"
)

type indentTest struct {
	indentLevel int
	tabs        int
	input       Line
	expected    Line
	err         error
}

// IndentTests are the test cases called by TestIndentEnvironments
var IndentTests = []indentTest{
	indentTest{
		indentLevel: 0,
		tabs:        1,
		input:       Line{0, "Hello!"},
		expected:    Line{0, "Hello!"},
		err:         nil},
	indentTest{
		indentLevel: 1,
		tabs:        1,
		input:       Line{0, "Hello!"},
		expected:    Line{0, "\tHello!"},
		err:         NoIndentErr{},
	},
	indentTest{
		indentLevel: 1,
		tabs:        1,
		input:       Line{0, "\tHello!"},
		expected:    Line{0, "\tHello!"},
		err:         nil,
	},
	indentTest{
		indentLevel: 2,
		tabs:        2,
		input:       Line{0, "\t\tHello!"},
		expected:    Line{0, "\t\t\t\tHello!"},
		err:         NoIndentErr{},
	},
	indentTest{
		indentLevel: 1,
		tabs:        2,
		input:       Line{0, "\t\tHello!"},
		expected:    Line{0, "\t\tHello!"},
		err:         nil,
	},
	indentTest{
		indentLevel: 0,
		tabs:        1,
		input:       Line{0, "\\begin"},
		expected:    Line{0, "\\begin"},
		err:         nil,
	},
	indentTest{
		indentLevel: 2,
		tabs:        1,
		input:       Line{0, "\\begin"},
		expected:    Line{0, "\t\t\\begin"},
		err:         NoIndentErr{}},
	indentTest{
		indentLevel: 0,
		tabs:        1,
		input:       Line{0, "\\end"},
		expected:    Line{0, "\\end"},
		err:         nil},
	indentTest{
		indentLevel: 0,
		tabs:        1,
		input:       Line{0, ""},
		expected:    Line{0, ""},
		err:         nil},
}

func TestIndentEnvironments(t *testing.T) {
	for _, test := range IndentTests {
		if output, _ := IndentEnvironments(test.indentLevel, test.tabs, test.input); output != test.expected {
			t.Errorf("Output %q not equal to expected %q. Types: %T %T", output, test.expected, output, test.expected)
		}
		if _, err := IndentEnvironments(test.indentLevel, test.tabs, test.input); !helper.CompareType(test.err, err) {
			t.Errorf("Output %q not equal to expected %q. Types: %T %T", err, test.err, err, test.err)
		}
	}
}
