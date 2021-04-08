package PostParser

import (
	"mdParser/Parse"
)

func removeEmptyLeaves(tree []Parse.ParseTree) []Parse.ParseTree {
	// Base case: empty tree
	if len(tree) == 0 {
		return tree
	}

	// Remove case: unit tree and reached empty text node
	if len(tree) == 1 &&
		len(tree[0].Children) <= 0 &&
		tree[0].Content == "" {
		return Parse.EmptyTree
	}

	// Unit tree and at expandable node  ->  Go deeper
	if len(tree) == 1 {
		tree[0].Children = removeEmptyLeaves(tree[0].Children)
		return tree
	}

	// Not a unit tree  ->  Recurse over all root nodes
	var newTree []Parse.ParseTree
	for _, node := range tree {
		newTree = append(newTree, removeEmptyLeaves(Parse.UnitTree(node))...)
	}
	return newTree
}
