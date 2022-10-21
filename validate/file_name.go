/*
Copyright © 2022 Filip Lindberg jakob.filip.lindberg@gmail.com
*/

package validate

import (
	"fmt"
	"regexp"
)

type BadFileNameError struct {
	Name string
}

func (e BadFileNameError) Error() string {
	return fmt.Sprintf("did not get a .tex file extension, got %s", e.Name)
}

func FileName(fileName string) error {
	if found, _ := regexp.MatchString("(.+.)(.tex|.bib|.tikz)$", fileName); !found {
		return BadFileNameError{fileName}
	}

	return nil
}
