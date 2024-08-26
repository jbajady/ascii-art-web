package Func

import (
	"os"
)

// read a banner file
func Readfile(banner string) ([]byte, error) {
	var contentfile []byte
	readbanner := banner + ".txt"
	contentfile, err := os.ReadFile(readbanner)
	return contentfile, err
}
