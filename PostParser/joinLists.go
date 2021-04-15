package PostParser

import (
	"mdParser/Parse"
)

func joinLists(tree []Parse.ParseTree) []Parse.ParseTree {
	// Base case: empty tree
	if len(tree) == 0 {
		return tree
	}

	// Unit tree and at expandable node  ->  Go deeper
	if len(tree) == 1 {
		tree[0].Children = joinLists(tree[0].Children)
		return tree
	}

	// Join case: Go through root nodes and join
	var joinedTree []Parse.ParseTree
	for i := 0; i < len(tree); i++ {
		node := tree[i]

		if node.TagName == Parse.UnorderedListTag || node.TagName == Parse.OrderedListTag {
			// Go through successive nodes and try to join
			for j := i+1; j < len(tree); j++ {
				if tree[j].TagName == node.TagName {
					// Join with previous node
					node.Children = append(node.Children, tree[j].Children...)
					i++ // Skip, since joining
				} else {
					break
				}
			}
		}

		joinedTree = append(joinedTree, node)
	}

	// Not a unit tree  ->  Recurse over all root nodes
	var newTree []Parse.ParseTree
	for _, node := range joinedTree {
		newTree = append(newTree, joinLists(Parse.UnitTree(node))...)
	}
	return newTree
}
