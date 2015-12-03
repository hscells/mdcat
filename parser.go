package main

import (
	"container/list"
	"strings"
)

// Parse takes an input string, separates it into an array split on new lines and parses each line using ParseLine
func Parse(input string) (ast list.List) {

	buffer := ""
	node := plaintext
	prevNode := plaintext
	startOfLine := true
	backtickCount := 0
	tildaCount := 0
	underlineCount := 0

	// ast := list.New()

	headingSize := 0

	for i := 0; i < len(input); i++ {
		char := input[i]

		// determing if this line is going to be a heading or blockquote
		if startOfLine {
			if node == plaintext {
				if char == '#' {
					node = atx
				} else if char == '>' {
					node = blockquote
				}
			}
		} else {
			buffer = ""
			node = plaintext
		}

		if char == '`' {
			backtickCount++
		} else if char == '~' {
			tildaCount++
		} else if char == '_' {
			underlineCount++
		} else if node == atx {
			if char == '#' {
				headingSize++
			} else {
				node = heading
				buffer = ""
			}
		} else if node == plaintext {
			if char == '*' {
				// add the previous text node before the `*`
				ast.PushBack(NewNodePair(node, buffer))
				prevNode = plaintext
				node = italic
				buffer = ""
			} else if tildaCount == 2 {
				tildaCount = 0
				if node == plaintext {
					ast.PushBack(NewNodePair(node, buffer))
					prevNode = plaintext
					node = strikethrough
					buffer = string(char)
				} else {
					ast.PushBack(NewNodePair(node, buffer))
					node = prevNode
					prevNode = strikethrough
					buffer = ""
				}
			} else if underlineCount == 2 {
				underlineCount = 0
				if node == plaintext {
					ast.PushBack(NewNodePair(node, buffer))
					prevNode = plaintext
					node = bold
					buffer = string(char)
				}
			} else {
				buffer += string(char)
			}
		} else if node == italic {
			if char == '*' {
				ast.PushBack(NewNodePair(node, buffer))
				node = prevNode
				buffer = ""
			} else {
				buffer += string(char)
			}
		} else if node == blockquote {
			ast.PushBack(NewNodePair(node, ""))
			node = plaintext
			buffer = ""
		} else if node == strikethrough {
			if tildaCount == 2 {
				// add the strike through text
				ast.PushBack(NewNodePair(node, buffer))
				tildaCount = 0
				node = prevNode
				buffer = ""
			} else {
				buffer += string(char)
			}
		} else if node == bold {
			if underlineCount == 2 {
				// add the bold text
				buffer += " "
				ast.PushBack(NewNodePair(node, buffer))
				underlineCount = 0
				node = prevNode
				buffer = ""
			} else {
				buffer += string(char)
			}
		}

		if node == heading {
			buffer += string(char)
		}

		// fmt.Println(tildaCount)

		if char == '\n' {
			if node == heading {
				headingNode := NewNodePair(heading, strings.TrimSpace(buffer))
				headingNode.Heading = headingSize
				ast.PushBack(headingNode)
				headingSize = 0
			} else {
				ast.PushBack(NewNodePair(node, strings.TrimSpace(buffer)))
			}
			ast.PushBack(NewNodePair(newline, ""))
			node = plaintext
			buffer = ""
		}
	}
	return ast
}
