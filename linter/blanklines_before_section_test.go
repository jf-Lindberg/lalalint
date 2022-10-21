/*
Copyright © 2022 Filip Lindberg jakob.filip.lindberg@gmail.com
*/

package linter

import (
	"github.com/jf-Lindberg/lalalint/helper"
	"testing"
)

type sectionTest struct {
	input            Line
	blankLinesCfg    int
	blankLinesBefore int
	expected         Line
	err              error
}

var sectionTests = []sectionTest{
	sectionTest{
		input:            CreateLine("no section here"),
		blankLinesCfg:    1,
		blankLinesBefore: 0,
		expected:         CreateLine("no section here"),
		err:              nil},
	sectionTest{
		input:            CreateLine("\\section"),
		blankLinesCfg:    1,
		blankLinesBefore: 0,
		expected:         CreateLine("\n\\section"),
		err:              NoBlankLineErr{}},
	sectionTest{
		input:            CreateLine("\\subsection"),
		blankLinesCfg:    5,
		blankLinesBefore: 4,
		expected:         CreateLine("\n\\subsection"),
		err:              NoBlankLineErr{}},
	sectionTest{
		input:            CreateLine("\\chapter"),
		blankLinesCfg:    2,
		blankLinesBefore: 0,
		expected:         CreateLine("\n\n\\chapter"),
		err:              NoBlankLineErr{}},
	sectionTest{
		input:            CreateLine("\\section"),
		blankLinesCfg:    1,
		blankLinesBefore: 1,
		expected:         CreateLine("\\section"),
		err:              nil},
	sectionTest{
		input:            CreateLine("section"),
		blankLinesCfg:    0,
		blankLinesBefore: 0,
		expected:         CreateLine("section"),
		err:              nil},
	sectionTest{
		input:            CreateLine(""),
		blankLinesCfg:    0,
		blankLinesBefore: 0,
		expected:         CreateLine(""),
		err:              nil},
}

func TestBlankLinesBeforeSection(t *testing.T) {
	for _, test := range sectionTests {
		if output, _ := BlankLinesBeforeSection(test.blankLinesCfg, test.blankLinesBefore, test.input); output != test.expected {
			t.Errorf("Output %q not equal to expected %q. Types: %T %T", output, test.expected, output, test.expected)
		}
		if _, err := BlankLinesBeforeSection(test.blankLinesCfg, test.blankLinesBefore, test.input); !helper.CompareType(test.err, err) {
			t.Errorf("Output %q not equal to expected %q. Types: %T %T", err, test.err, err, test.err)
		}
	}
}
