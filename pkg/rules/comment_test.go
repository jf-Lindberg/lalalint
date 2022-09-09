/*
Copyright © 2022 Filip Lindberg jakob.filip.lindberg@gmail.com
*/

package rules

import "testing"

type commentTest struct {
	input    string
	expected string
}

var commentTests = []commentTest{
	commentTest{input: "%this is wrong", expected: "% this is wrong"},
	commentTest{input: "% this is right", expected: "% this is right"},
	commentTest{input: `\%this is escaped`, expected: `\%this is escaped`},
	commentTest{input: "no comment here", expected: "no comment here"},
	commentTest{input: "1233213", expected: "1233213"},
	commentTest{input: "inline %comment wrong", expected: "inline % comment wrong"},
	commentTest{input: `inline \%comment escaped`, expected: `inline \%comment escaped`},
}

func TestComment(t *testing.T) {
	for _, test := range commentTests {
		if output := LintComment(test.input); output != test.expected {
			t.Errorf("Output %q not equal to expected %q. Types: %T %T", output, test.expected, output, test.expected)
		}
	}
}
