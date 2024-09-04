package ascii

import (
	"fmt"
)

func PrintArt(string, []string, map[rune]([]string)) {
	// we start by checking if the user input only contains literal newlines
	// if so we print newlines accordingly
	if ContainsOnly(input) {
		for i := 0; i < len(input)/2; i++ {
			fmt.Println()
		}
		return
	}
	for _, line := range inputsplit {
		// also if there's empty strings resulting from the spliting we print a newline
		if line == "" {
			fmt.Println()
			continue
		}
		// now the printing, first loop controls the lines, second prints the characters
		for i := 0; i < 8; i++ {
			for j := 0; j < len(line); j++ {
				inputrune := rune(line[j])
				fmt.Print(Replace[inputrune][i])
			}
		// after each line is printed we print a newline
			fmt.Println()
		}
	}
}
