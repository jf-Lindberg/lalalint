/*
Copyright © 2022 Filip Lindberg fili21@student.bth.se
*/

package linter

import (
	"github.com/jf-Lindberg/lalalint/helper"
	"testing"
)

type newLineAfterSentenceTest struct {
	input    Line
	expected Line
	err      error
}

// IndentTests are the test cases called by TestIndentEnvironments
var NewLineAfterSentenceTests = []newLineAfterSentenceTest{
	newLineAfterSentenceTest{
		input:    Line{0, "Hello! More words"},
		expected: Line{0, "Hello!\nMore words"},
		err:      NewlineAfterSentenceErr{}},
	newLineAfterSentenceTest{
		input:    Line{0, "Hello!"},
		expected: Line{0, "Hello!"},
		err:      nil},
	newLineAfterSentenceTest{
		input:    Line{0, "Hello ! More words"},
		expected: Line{0, "Hello ! More words"},
		err:      nil},
	newLineAfterSentenceTest{
		input:    Line{0, "This is a sentence? And more."},
		expected: Line{0, "This is a sentence?\nAnd more."},
		err:      NewlineAfterSentenceErr{}},
	newLineAfterSentenceTest{
		input:    Line{0, "This is a sentence! And more."},
		expected: Line{0, "This is a sentence!\nAnd more."},
		err:      NewlineAfterSentenceErr{}},
	newLineAfterSentenceTest{
		input:    Line{0, "%This is a comment! And more."},
		expected: Line{0, "%This is a comment!\n%And more."},
		err:      NewlineAfterSentenceErr{}},
	newLineAfterSentenceTest{
		input:    Line{0, "This is a weird . sentence"},
		expected: Line{0, "This is a weird . sentence"},
		err:      nil},
	newLineAfterSentenceTest{
		input:    Line{0, "example.com"},
		expected: Line{0, "example.com"},
		err:      nil},
	newLineAfterSentenceTest{
		input:    Line{0, "3.14! Otherwise known as pi"},
		expected: Line{0, "3.14!\nOtherwise known as pi"},
		err:      NewlineAfterSentenceErr{}},
	newLineAfterSentenceTest{
		input:    Line{0, ""},
		expected: Line{0, ""},
		err:      nil},
}

func TestNewlineAfterSentence(t *testing.T) {
	for _, test := range NewLineAfterSentenceTests {
		if output, _ := NewlineAfterSentence(test.input); output != test.expected {
			t.Errorf("Output %q not equal to expected %q. Types: %T %T", output, test.expected, output, test.expected)
		}
		if _, err := NewlineAfterSentence(test.input); !helper.CompareType(test.err, err) {
			t.Errorf("Output %q not equal to expected %q. Types: %T %T", err, test.err, err, test.err)
		}
	}
}
