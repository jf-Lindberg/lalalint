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

var commentTests = []commentTest{
	commentTest{input: createLine("%this is wrong"), expected: createLine("% this is wrong"), err: BadCommentError{}},
	commentTest{input: createLine("% this is right"), expected: createLine("% this is right"), err: nil},
	commentTest{input: createLine(`\%this is escaped`), expected: createLine(`\%this is escaped`), err: nil},
	commentTest{input: createLine("no comment here"), expected: createLine("no comment here"), err: nil},
	commentTest{input: createLine("1233213"), expected: createLine("1233213"), err: nil},
	commentTest{input: createLine("inline %comment wrong"), expected: createLine("inline % comment wrong"), err: BadCommentError{}},
	commentTest{input: createLine(`inline \%comment escaped`), expected: createLine(`inline \%comment escaped`), err: nil},
}

func createLine(input string) Line {
	return Line{0, input}
}

func TestComment(t *testing.T) {
	for _, test := range commentTests {
		if output, _ := LintComment(test.input, "%"); output != test.expected {
			t.Errorf("Output %q not equal to expected %q. Types: %T %T", output, test.expected, output, test.expected)
		}
		if _, err := LintComment(test.input, "%"); !helper.CompareType(test.err, err) {
			t.Errorf("Output %q not equal to expected %q. Types: %T %T", err, test.err, err, test.err)
		}
	}
}
