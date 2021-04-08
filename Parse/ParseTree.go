package Parse

type ParseTree struct {
	TagName  Tag
	Children []ParseTree
	Content  string
}

// Not to mistake with a leaf node, a unit tree has only one node at its root
func UnitTree(node ParseTree) []ParseTree {
	return []ParseTree{node}
}

// Leaf node
func Text(text string) ParseTree {
	return ParseTree{
		TagName: TextTag,
		Content: text,
	}
}

// Unexpanded node
func Raw(text string) ParseTree {
	return ParseTree{
		TagName: RawTag,
		Content: text,
	}
}

// List containing one unexpanded node
func RawChild(text string) []ParseTree{
	return UnitTree(Raw(text))
}
