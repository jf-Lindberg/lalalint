/*
Copyright © 2022 Filip Lindberg jakob.filip.lindberg@gmail.com
*/

package linter

import (
	"github.com/jf-Lindberg/lalalint/helper"
	"testing"
)

type commentTest struct {
	input    Line
	expected Line
	err      error
}

// commentTests are the test cases called by TestSpaceAfterComment
var commentTests = []commentTest{
	commentTest{
		input:    CreateLine("%this is wrong"),
		expected: CreateLine("% this is wrong"),
		err:      BadCommentError{}},
	commentTest{
		input:    CreateLine("% this is right"),
		expected: CreateLine("% this is right"),
		err:      nil},
	commentTest{
		input:    CreateLine(`\%this is escaped`),
		expected: CreateLine(`\%this is escaped`),
		err:      nil},
	commentTest{
		input:    CreateLine("no comment here"),
		expected: CreateLine("no comment here"),
		err:      nil},
	commentTest{
		input:    CreateLine("1233213"),
		expected: CreateLine("1233213"),
		err:      nil},
	commentTest{
		input:    CreateLine("inline %comment wrong"),
		expected: CreateLine("inline % comment wrong"),
		err:      BadCommentError{}},
	commentTest{
		input:    CreateLine(`inline \%comment escaped`),
		expected: CreateLine(`inline \%comment escaped`),
		err:      nil},
	commentTest{
		input:    CreateLine(""),
		expected: CreateLine(""),
		err:      nil},
}

func TestSpaceAfterComment(t *testing.T) {
	for _, test := range commentTests {
		if output, _ := SpaceAfterComment(test.input); output != test.expected {
			t.Errorf("Output %q not equal to expected %q. Types: %T %T", output, test.expected, output, test.expected)
		}
		if _, err := SpaceAfterComment(test.input); !helper.CompareType(test.err, err) {
			t.Errorf("Output %q not equal to expected %q. Types: %T %T", err, test.err, err, test.err)
		}
	}
}
