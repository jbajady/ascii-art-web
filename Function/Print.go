package Func

// print the sscii-art
func Print(inputtext []string, matrix [][]string) string {
	resulte := ""
	for j := 0; j < len(inputtext); j++ {
		if inputtext[j] == "" { // replece any umpty string with newline
			resulte += "\n"
			continue

		}
		for i := 0; i < 8; i++ { // print 8 line
			for k := 0; k < len(inputtext[j]); k++ {
				resulte += (matrix[int(inputtext[j][k]-32)][i]) // print the asci-art
			}
			resulte += "\n"
		}

	}
	return resulte
}
