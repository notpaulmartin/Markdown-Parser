package Rules

import (
	"mdParser/Parse"
	"mdParser/Parse/RuleParser"
)

var (
	Headings = RuleParser.Precedence{Order: []Parse.Applyable{
		&Heading1,
		&Heading2,
		&Heading3,
		&Heading4,
		&Heading5,
		&Heading6,
	}}

	Heading1 = RuleParser.Rule{ApplyFunc: extractRegex(Parse.H1Tag, "^# (.*)$")}
	Heading2 = RuleParser.Rule{ApplyFunc: extractRegex(Parse.H2Tag, "^## (.*)$")}
	Heading3 = RuleParser.Rule{ApplyFunc: extractRegex(Parse.H3Tag, "^### (.*)$")}
	Heading4 = RuleParser.Rule{ApplyFunc: extractRegex(Parse.H4Tag, "^#### (.*)$")}
	Heading5 = RuleParser.Rule{ApplyFunc: extractRegex(Parse.H5Tag, "^##### (.*)$")}
	Heading6 = RuleParser.Rule{ApplyFunc: extractRegex(Parse.H6Tag, "^###### (.*)$")}
)
