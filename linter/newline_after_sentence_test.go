/*
Copyright © 2022 Filip Lindberg jakob.filip.lindberg@gmail.com
*/

package linter

import (
	"testing"
)

type newlineAfterSentenceTest struct {
	input    Line
	expected []NewLine
}

var NewlineAfterSentenceTests = []newlineAfterSentenceTest{
	{input: Line{0, "Hello!"}, expected: []NewLine{{Line{0, "Hello!"}, nil}}},
}

func TestNewlineAfterSentence(t *testing.T) {
	for _, test := range NewlineAfterSentenceTests {
		if output := NewlineAfterSentence(test.input); output[0].line.Content != test.expected[0].line.Content {
			t.Errorf("Output %v not equal to expected %v. Types: %T %T", output, test.expected, output, test.expected)
		}
	}
}
