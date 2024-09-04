package ascii

import (
	"log"
	"os"
	"strings"
)

var Banner string

func BannerManagement() string {
	args := os.Args[1:]
	// Usage instructions if no arguments or more than 2 arguments are provided
	if len(args) == 0 || len(args) > 2 {
		log.Fatal("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
	}
	// the default styling will be standard.txt unless the user chooses otherwise
	Banner = "standard.txt"
	if len(args) == 2 {
		Banner = args[1]
	} else if len(args) == 1 {
		args = append(args, Banner)
	}
	// we add a .txt extension to the banner if it doesn't already exist
	//this way weather the user adds it or not it won't make a difference
	if !strings.HasSuffix(args[1], ".txt") {
		args[1] += ".txt"
	}
	//we inform the user of the available styles if he doesn't choose a proper one
	if args[1] != "standard.txt" && args[1] != "thinkertoy.txt" && args[1] != "shadow.txt" {
		log.Fatal("Usage: this style is unavailable \nPlease choose one of the available styles \n1 : standard \n2 : thinkertoy \n3 : shadow")
	}

	return Banner
}
