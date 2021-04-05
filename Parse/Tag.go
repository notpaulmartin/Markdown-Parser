package Parse

type Tag struct {
	Name     string
	Content  string
}

func Text(text string) Tag {
	return Tag{
		Name:    "text",
		Content: text,
	}
}

func Raw(text string) Tag {
	return Tag{
		Name:    "raw",
		Content: text,
	}
}
