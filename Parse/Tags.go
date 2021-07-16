package Parse

type Tag string

const (
	RawTag  Tag = "RAW"
	TextTag Tag = "TEXT"

	EscapeTag Tag = "ESCAPE"

	HtmlTagTag Tag = "HTML_TAG"

	SectionTag Tag = "section"

	H1Tag Tag = "h1"
	H2Tag Tag = "h2"
	H3Tag Tag = "h3"
	H4Tag Tag = "h4"
	H5Tag Tag = "h5"
	H6Tag Tag = "h6"

	ParagraphTag Tag = "p"

	BoldTag          Tag = "strong"
	ItalicsTag       Tag = "em"
	LinkTag          Tag = "a"
	ImgTag           Tag = "img"
	StrikethroughTag Tag = "strike"

	CheckboxTrueTag  Tag = "checkboxTrue"  // Checked
	CheckboxFalseTag Tag = "checkboxFalse" // Unchecked
	OrderedListTag   Tag = "ol"
	UnorderedListTag Tag = "ul"
	ListItemTag      Tag = "li"
)
