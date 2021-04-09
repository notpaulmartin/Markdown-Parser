package Compiler

import (
	"fmt"
	"mdParser/Parse"
	"strings"
)

func ToHtml(tree []Parse.ParseTree) string {
	strBuilder := &strings.Builder{}
	strBuilder = toHtml(tree, strBuilder)
	return strBuilder.String()
}

// Recursive helper-function for ToHtml
func toHtml(tree []Parse.ParseTree, strBuilder *strings.Builder) *strings.Builder {
	// Base case: empty tree
	if len(tree) == 0 {
		return strBuilder
	}

	// Recursive case: Multiple root nodes  ->  Recurse over root nodes
	if len(tree) > 1 {
		for _, subtree := range tree {
			strBuilder = toHtml(Parse.UnitTree(subtree), strBuilder)
		}

		return strBuilder
	}

	// Is Unit-tree
	node := tree[0]

	if node.TagName == Parse.TextTag {
		strBuilder.WriteString(node.Content)
		return strBuilder
	}

	strBuilder.WriteString(openingHtmlTag(node))
	toHtml(node.Children, strBuilder)
	strBuilder.WriteString(closingHtmlTag(node))

	return strBuilder
}

func openingHtmlTag(node Parse.ParseTree) string {
	switch node.TagName {
	case Parse.LinkTag:
		return fmt.Sprintf("<a href=\"%s\">", node.Content)
	case Parse.ImgTag:
		return fmt.Sprintf("<img src=\"%s\">", node.Content)
	default:
		return fmt.Sprintf("<%s>", node.TagName)
	}

	// TODO: Checkbox, List
}

func closingHtmlTag(node Parse.ParseTree) string {
	switch node.TagName {
	default:
		return fmt.Sprintf("</%s>", node.TagName)
	}

	// TODO: Checkbox, List
}
