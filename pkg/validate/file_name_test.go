/*
Copyright © 2022 Filip Lindberg jakob.filip.lindberg@gmail.com
*/

package validate

import "testing"

type nameTest struct {
	name     string
	expected error
}

var nameTests = []nameTest{
	nameTest{name: "test.t", expected: BadFileNameError{Name: "test.t"}},
	nameTest{name: "12313312312", expected: BadFileNameError{Name: "12313312312"}},
	nameTest{name: "", expected: BadFileNameError{Name: ""}},
	nameTest{name: "tex.j", expected: BadFileNameError{Name: "tex.j"}},
	nameTest{name: "</%", expected: BadFileNameError{Name: "</%"}},
	nameTest{name: "test.te.x", expected: BadFileNameError{Name: "test.te.x"}},
	nameTest{name: ".tex", expected: BadFileNameError{Name: ".tex"}},
	nameTest{name: "test.tex", expected: nil},
	nameTest{name: "TEsTinGNIINENin.tex", expected: nil},
	nameTest{name: "388hej.tex", expected: nil},
}

func Test_FileName(t *testing.T) {
	for _, test := range nameTests {
		if output := FileName(test.name); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
