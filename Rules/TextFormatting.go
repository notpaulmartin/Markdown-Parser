package Rules

import (
	"mdParser/Parse"
	"mdParser/Parse/RuleParser"
)

// Text formatting (bold, italics, etc.)
var (
	TextFormatting = RuleParser.Precedence{Order: []Parse.Applyable{
		// &BoldItalics,
		&Bold,
		&Italics,
	}}

	Bold    = RuleParser.Rule{
		TagName: Parse.BoldTag,
		ApplyFunc: applyRegexInText(Parse.BoldTag, `\*\*(.*)\*\*`),
	}
	Italics = RuleParser.Rule{
		TagName: Parse.ItalicsTag,
		ApplyFunc: applyRegexInText(Parse.ItalicsTag, `\*(.*)\*`),
	}
	// /*?*/ BoldItalics = Rule{applyRegexInText(Parse.ItalicsTag, `\*\*\*(.*)\*\*\*`)}
)
