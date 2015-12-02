package main

import (
	"regexp"
	"strings"
)

// ParseLine takes one line of markdown and decides how to render it
func ParseLine(input string) {
	if len(input) == 0 {
		return
	}
	if input[0] == '#' {
		level := 1
		for i := 1; i < len(input); i++ {
			if input[i] != '#' {
				break
			}
			level++
		}
		PrintAtx(strings.TrimSpace(input[level:]), level)
	} else if input[0] == '>' {
		PrintBlockQuote(input[1:])
	} else {
		Print(input)
	}
}

// Parse takes an input string, separates it into an array split on new lines and parses each line using ParseLine
func Parse(input string) {
	arr := strings.Split(input, "\n")
	previousLine := ""
	for i := 0; i < len(arr); i++ {
		// grab the current line being looked at
		line := arr[i]

		// trim the start and end from the string
		line = strings.TrimSpace(line)

		setextH1, _ := regexp.MatchString("=+$", line)
		setextH2, _ := regexp.MatchString("-+$", line)

		// parse a single line of input
		if setextH1 {
			PrintSetext(previousLine, 1)
			line = ""
		} else if setextH2 {
			PrintSetext(previousLine, 2)
			line = ""
		} else {
			ParseLine(previousLine)
		}
		previousLine = line
	}
}
