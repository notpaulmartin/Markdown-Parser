package PostParser

import (
	"mdParser/Parse"
	"reflect"
)

func Clean(tree []Parse.ParseTree) []Parse.ParseTree {
	var prevTree []Parse.ParseTree

	for !reflect.DeepEqual(prevTree, tree) {
		prevTree = make([]Parse.ParseTree, len(tree))
		copy(prevTree, tree)

		tree = removeEmptyLeaves(tree)
	}

	return tree
}
