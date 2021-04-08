package Rules

import (
	"mdParser/Parse"
	"mdParser/Parse/RuleParser"
)

var Formatters = RuleParser.Precedence{Order: []Parse.Applyable{
	&Specials,
	&TextFormatting,
	&Text,
}}

// Text formatting (bold, italics, etc.)
var (
	TextFormatting = RuleParser.Precedence{Order: []Parse.Applyable{
		// &BoldItalics,
		&Bold,
		&Italics,
	}}

	Bold    = RuleParser.Rule{ApplyFunc: applyRegexInText(Parse.BoldTag, `\*\*(.*)\*\*`)}
	Italics = RuleParser.Rule{ApplyFunc: applyRegexInText(Parse.ItalicsTag, `\*(.*)\*`)}
	// /*?*/ BoldItalics = Rule{applyRegexInText(Parse.ItalicsTag, `\*\*\*(.*)\*\*\*`)}
)

// Specials (links, pics, etc.)
var (
	Specials = RuleParser.Precedence{Order: []Parse.Applyable{
		&Pic,
		&Link,
	}}

	Pic  = RuleParser.Rule{applyRegexInText(Parse.PicTag, `![.*](.*)`)}
	Link = RuleParser.Rule{applyRegexInText(Parse.LinkTag, `[.*](.*)`)}
)
