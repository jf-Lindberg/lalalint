/*
Copyright © 2022 Filip Lindberg jakob.filip.lindberg@gmail.com
*/

package cmd

import (
	"errors"
	"fmt"
	"github.com/jf-Lindberg/lalalint/pkg/validate"
)

import (
	"bytes"
	"testing"
)

type argTest struct {
	arg      string
	expected error
}

var argTests = []argTest{
	argTest{arg: "test.t", expected: validate.BadFileNameError{Name: "test.t"}},
	argTest{arg: "12313312312", expected: validate.BadFileNameError{Name: "12313312312"}},
	argTest{arg: "", expected: validate.BadFileNameError{Name: ""}},
	argTest{arg: "tex.j", expected: validate.BadFileNameError{Name: "tex.j"}},
	argTest{arg: "</%", expected: validate.BadFileNameError{Name: "</%"}},
	argTest{arg: "test.te.x", expected: validate.BadFileNameError{Name: "test.te.x"}},
	argTest{arg: ".tex", expected: validate.BadFileNameError{Name: ".tex"}},
	argTest{arg: "test.tex", expected: nil},
	argTest{arg: "TEsTinGNIINENin.tex", expected: nil},
	argTest{arg: "388hej.tex", expected: nil},
}

func TestWithArg(t *testing.T) {
	cmd := rootCmd
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	for _, test := range argTests {
		b.Reset()
		cmd.SetArgs([]string{test.arg})
		if output := cmd.Execute(); output != test.expected {
			t.Errorf("Output %q not equal to expected %q. Types: %T %T", output, test.expected, output, test.expected)
		}
	}
}

func TestNoArgs(t *testing.T) {
	cmd := rootCmd
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{})
	expected := errors.New("accepts 1 arg(s), received 0")
	if err := cmd.Execute(); err.Error() != expected.Error() {
		t.Errorf(fmt.Sprintf("Output >%s< not equal to expected >%s<. Types: %T %T", err, expected, err, expected))
	}
}
