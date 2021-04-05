package Rules

var (
	Headings = Precedence{order: []Rule{
		Heading1,
		Heading2,
		Heading3,
		Heading4,
		Heading5,
		Heading6,
	}}

	Heading1 = Rule{Apply: applyRegex("h1", "^# (.*)$")}
	Heading2 = Rule{Apply: applyRegex("h2", "^## (.*)$")}
	Heading3 = Rule{Apply: applyRegex("h3", "^### (.*)$")}
	Heading4 = Rule{Apply: applyRegex("h4", "^#### (.*)$")}
	Heading5 = Rule{Apply: applyRegex("h5", "^##### (.*)$")}
	Heading6 = Rule{Apply: applyRegex("h6", "^###### (.*)$")}
)
