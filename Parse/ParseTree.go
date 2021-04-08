package Parse

type ParseTree struct {
	TagName  Tag
	Children []ParseTree
	Content  string
}
