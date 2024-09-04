package main

import (
	"ascii/ascii"
)

func main() {
	// we get the banner style and input text from user input
	banner := ascii.BannerManagement()
	input, inputSplit := ascii.InputManagement()
	// we read the ASCII art characters from the chosen banner file
	ascii.ReadText(banner)

	// we populate the map with ASCII characters
	replaceMap := ascii.Populate()

	// and finally we print the resulting ascii art

	ascii.PrintArt(input, inputSplit, replaceMap)
}
