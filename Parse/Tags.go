package Parse

type Tag string

const (
	RawTag  Tag = "RAW"
	TextTag Tag = "TEXT"

	H1Tag Tag = "h1"
	H2Tag Tag = "h2"
	H3Tag Tag = "h3"
	H4Tag Tag = "h4"
	H5Tag Tag = "h5"
	H6Tag Tag = "h6"

	BoldTag          Tag = "b"
	ItalicsTag       Tag = "i"
	StrikethroughTag Tag = "strike"

	CheckboxTrueTag  Tag = "checkboxTrue"  // Checked
	CheckboxFalseTag Tag = "checkboxFalse" // Unchecked
	OrderedListTag   Tag = "orderedList"
	UnorderedListTag Tag = "unorderedList"
)
