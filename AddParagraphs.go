package main

import (
	"mdParser/Parse"
	"mdParser/Rules"
)


// Similar to RuleParser.RecursiveApply, going through tree, but only converting
// uninterpreted RAW nodes at the root (or in html tags) to paragraphs
func AddParagraphs(tree []Parse.ParseTree) []Parse.ParseTree {
	// Base case: empty tree
	if len(tree) == 0 {
		return tree
	}

	// Base case: unit tree and reached RAW
	if len(tree) == 1 && tree[0].TagName == Parse.RawTag {
		success, parse := Rules.Paragraph.Apply(tree[0].Content)
		if !success {
			// If can't make paragraph, leave as is
			return tree
		}
		return parse
	}

	// Unit tree and at expandable node  ->  Go deeper
	if len(tree) == 1 && tree[0].TagName == Parse.HtmlTagTag {
		tree[0].Children = AddParagraphs(tree[0].Children)
		return tree
	}

	// Base case: unit tree and reached leaf
	if len(tree) == 1 {
		return tree
	}

	// Not a unit tree  ->  Recurse over all root nodes
	var newTree []Parse.ParseTree
	for _, node := range tree {
		newTree = append(newTree, AddParagraphs(Parse.UnitTree(node))...)
	}
	return newTree
}

