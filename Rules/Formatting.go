package Rules

var (
	Bold = Rule{applyRegex2("bold", `\*\*(.*)\*\*`)}
	Italics = Rule{applyRegex2("it", `\*(.*)\*`)}
	/*?*/ BoldItalics = Rule{applyRegex2("bold,it", `\*\*\*(.*)\*\*\*`)}

	UnorderedListItem = Rule{applyRegex("UnorderedListItem", `- (.*)`)}
	OrderedListItem = Rule{applyRegex("OrderedListItem", `\d+\. (.*)`)}
)
