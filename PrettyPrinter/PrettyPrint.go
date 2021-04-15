package PrettyPrinter

import (
	"fmt"
	"mdParser/Parse"
	"strings"
)

const indentStr = "|   "

func PrettyPrint(tree []Parse.ParseTree) {
	prettyPrint_helper(tree, 0)
}

func prettyPrint_helper(tree []Parse.ParseTree, indent int) {
	if len(tree) == 0 {
		return
	}

	// If tree has multiple nodes, recurse over nodes with an incremented indent
	if len(tree) > 1 {
		for _, subtree := range tree {
			prettyPrint_helper(Parse.UnitTree(subtree), indent)
		}
		return
	}

	node := tree[0]

	// Print indent
	fmt.Printf("%v", strings.Repeat(indentStr, indent))
	// Print tag
	fmt.Printf("<%v>: ", node.TagName)
	// Print content
	if node.Content != "" {
		fmt.Printf("%q", node.Content)
	}
	fmt.Println()

	// Recurse over children
	prettyPrint_helper(node.Children, indent+1)
}
