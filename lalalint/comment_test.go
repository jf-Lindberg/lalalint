/*
Copyright © 2022 Filip Lindberg jakob.filip.lindberg@gmail.com
*/

package lalalint

import (
	"testing"
)

type commentTest struct {
	input    Line
	expected Line
}

var commentTests = []commentTest{
	commentTest{input: createLine("%this is wrong"), expected: createLine("% this is wrong")},
	commentTest{input: createLine("% this is right"), expected: createLine("% this is right")},
	commentTest{input: createLine(`\%this is escaped`), expected: createLine(`\%this is escaped`)},
	commentTest{input: createLine("no comment here"), expected: createLine("no comment here")},
	commentTest{input: createLine("1233213"), expected: createLine("1233213")},
	commentTest{input: createLine("inline %comment wrong"), expected: createLine("inline % comment wrong")},
	commentTest{input: createLine(`inline \%comment escaped`), expected: createLine(`inline \%comment escaped`)},
}

func createLine(input string) Line {
	return Line{0, input}
}

func TestComment(t *testing.T) {
	for _, test := range commentTests {
		if output, _ := LintComment(test.input, "%"); output != test.expected {
			t.Errorf("Output %q not equal to expected %q. Types: %T %T", output, test.expected, output, test.expected)
		}
	}
}

//func TestError() {
//
//}
