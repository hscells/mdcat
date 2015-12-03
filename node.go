package main

const (
	atx = 1 << iota
	se
	heading
	blockquote
	blockquotetext
	italic
	bold
	underline
	strikethrough
	code
	plaintext
	newline
)

// Node is a tuple which contains the id of the node and the text content of that node
type Node struct {
	Type    int
	Content string

	Heading int
}

// NewNodePair creates a new node pair object
func NewNodePair(nodeType int, nodeText string) (nodePair Node) {
	nodePair = Node{}
	nodePair.Type = nodeType
	nodePair.Content = nodeText
	return
}
