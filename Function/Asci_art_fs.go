package Func

import (
	"strings"
)

// ascii-art-fs
func Ascii_art_fs(inputtext string, banner string) string {
	var word string

	inputsep := strings.Split(inputtext, "\r\n")
	if strings.ReplaceAll(inputtext, "\r\n", "") == "" && len(inputtext) > 1 {
		inputsep = inputsep[1:]
	}
	contentfile := Readfile(banner)
	matrix := Split((string(contentfile)), banner)
	word = Print(inputsep, matrix)
	return word
}
