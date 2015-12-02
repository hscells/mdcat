package main

import (
	"fmt"
	"github.com/fatih/color"
)

// PrintHeading prints an arbitraty heading to the screen, the textContent is what is to be displayed and the level is
// the level at which the heading is formatted to
func PrintHeading(textContent string, level int) {
	fmt.Print("  ")
	if level > 6 {
		level = 6
	}
	if level < 1 {
		level = 1
	}
	for i := 1; i < level+1; i++ {
		fmt.Print("  ")
	}
	c := color.New(color.Underline)
	c.Println(textContent)
	fmt.Println()
}

// PrintSetext prints a setext which is indicated in markdown using
//
// text content
// ============
//
// or
//
// text content
// ------------
//
// level indicates the heading level (1 - 2)
//
func PrintSetext(textContent string, level int) {
	PrintHeading(textContent, level)
}

// PrintAtx prints an atx heading, which is indicated in markdown using
//
// #text content
// up to
// ######text content
//
// level is the number of pound signs used (from 1 - 6)
//
func PrintAtx(textContent string, level int) {
	PrintHeading(textContent, level)
}

// PrintBlockQuote takes a string and styles it as a blockquote
func PrintBlockQuote(textContent string) {
	fmt.Print("  ")
	c := color.New(color.Italic).Add(color.Faint)
	c.Println(textContent)
}

// PrintItalic takes a string and prints it italicised
func PrintItalic(textContent string) {
	c := color.New(color.Italic)
	c.Print(textContent)
}

// Print simply prints a text content node to the screen
func Print(textContent string) {
	fmt.Println(textContent)
}
