package main

import (
	"container/list"
	"fmt"
	"github.com/fatih/color"
)

// Render takes an ast list and renders it to the screen
func Render(ast list.List) {
	for e := ast.Front(); e != nil; e = e.Next() {
		node, ok := e.Value.(Node)
		if !ok {
			panic(ok)
		}
		if node.Type == newline {
			fmt.Println()
		} else {
			if node.Type == heading {
				PrintHeading(node.Content, node.Heading)
			} else if node.Type == blockquote {
				fmt.Print("\t")
			} else if node.Type == italic {
				PrintItalic(node.Content)
			} else if node.Type == strikethrough {
				PrintStrikethrough(node.Content)
			} else if node.Type == bold {
				PrintBold(node.Content)
			} else {
				Print(node.Content)
			}
		}
	}
}

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
	c.Print(textContent)
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

// PrintBold takes a string and prints it italicised
func PrintBold(textContent string) {
	c := color.New(color.Bold)
	c.Print(textContent)
}

// PrintStrikethrough takes a string and prints it italicised
func PrintStrikethrough(textContent string) {
	c := color.New(color.CrossedOut)
	c.Print("~~" + textContent + "~~")
}

// Print simply prints a text content node to the screen
func Print(textContent string) {
	fmt.Print(textContent)
}
