package Rules

import "mdParser/Parse"

/* TODO:
 *  - Multi-line rules
 *  	-> Code-blocks (tab / ```)
 *  -
*/

var All = Precedence{[]Applyable{
	&Extractors,
	&Formatters,
	&Text,
}}

var Extractors = Precedence{[]Applyable{
	&Headings,
	&Lists,
}}

var Formatters = Precedence{[]Applyable{
	&TextFormatting,
	&Text,
}}

/// Text - If everything else fails
var Text = Rule{ApplyFunc: func(s string) (bool, []Parse.ParseTree) {
	return true, []Parse.ParseTree{Parse.Text(s)}
}}

var (
	TextFormatting = Precedence{[]Applyable{
		// &BoldItalics,
		&Bold,
		&Italics,
	}}

	Bold    = Rule{applyRegexInText(Parse.BoldTag, `\*\*(.*)\*\*`)}
	Italics = Rule{applyRegexInText(Parse.ItalicsTag, `\*(.*)\*`)}
	// /*?*/ BoldItalics = Rule{applyRegexInText(Parse.ItalicsTag, `\*\*\*(.*)\*\*\*`)}
)

/// Extractors

var (
	Headings = Precedence{order: []Applyable{
		&Heading1,
		&Heading2,
		&Heading3,
		&Heading4,
		&Heading5,
		&Heading6,
	}}

	Heading1 = Rule{ApplyFunc: extractRegex(Parse.H1Tag, "^# (.*)$")}
	Heading2 = Rule{ApplyFunc: extractRegex(Parse.H2Tag, "^## (.*)$")}
	Heading3 = Rule{ApplyFunc: extractRegex(Parse.H3Tag, "^### (.*)$")}
	Heading4 = Rule{ApplyFunc: extractRegex(Parse.H4Tag, "^#### (.*)$")}
	Heading5 = Rule{ApplyFunc: extractRegex(Parse.H5Tag, "^##### (.*)$")}
	Heading6 = Rule{ApplyFunc: extractRegex(Parse.H6Tag, "^###### (.*)$")}
)

var (
	Lists = Precedence{[]Applyable{
		&CheckboxChecked,
		&CheckboxUnchecked,
		&UnorderedListItem,
		&OrderedListItem,
	}}

	CheckboxChecked = Rule{ApplyFunc: extractRegex(Parse.CheckboxTrueTag, `- \[[x|X]\] (.*)`)}
	CheckboxUnchecked = Rule{ApplyFunc: extractRegex(Parse.CheckboxFalseTag, `- \[ ?\] (.*)`)}

	UnorderedListItem = Rule{ApplyFunc: extractRegex(Parse.UnorderedListTag, `- (.*)`)}
	OrderedListItem   = Rule{ApplyFunc: extractRegex(Parse.OrderedListTag, `\d+\. (.*)`)}
)
