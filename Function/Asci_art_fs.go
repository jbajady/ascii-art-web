package Func

import (
	"strings"
)

// ascii-art-fs
func Ascii_art_fs(inputtext string, banner string) (string, error) {
	var resulte string

	inputsep := strings.Split(inputtext, "\r\n")
	if strings.ReplaceAll(inputtext, "\r\n", "") == "" && len(inputtext) > 1 {
		inputsep = inputsep[1:]
	}
	contentfile, err := Readfile(banner)
	if err != nil {
		return resulte, err
	}
	matrix := Split((string(contentfile)), banner)
	resulte = Print(inputsep, matrix)
	return resulte, err
}
