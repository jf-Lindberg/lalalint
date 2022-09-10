package rules

func Caller(line string) (string, error) {
	output := LintComment(line)
	// call other with output of output
	return output, nil
}
