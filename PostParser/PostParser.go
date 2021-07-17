package PostParser

import (
	"github.com/notpaulmartin/mdParser/Parse"
	"reflect"
)

// TODO:
//  - Join adjacent texts into one

// main
func Clean(tree []Parse.ParseTree) []Parse.ParseTree {
	var prevTree []Parse.ParseTree

	// Do until clean
	for !reflect.DeepEqual(prevTree, tree) {
		prevTree = make([]Parse.ParseTree, len(tree))
		copy(prevTree, tree)

		// Cleaning functions
		tree = removeEmptyLeaves(tree)
		tree = joinLists(tree)
	}

	return tree
}
