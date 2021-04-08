package Rules

import (
	"mdParser/Parse"
	"mdParser/Parse/RuleParser"
)

var (
	Lists = RuleParser.Precedence{Order: []Parse.Applyable{
		&CheckboxChecked,
		&CheckboxUnchecked,
		&UnorderedListItem,
		&OrderedListItem,
	}}

	CheckboxChecked = RuleParser.Rule{ApplyFunc: extractRegex(Parse.CheckboxTrueTag, `- \[[x|X]\] (.*)`)}
	CheckboxUnchecked = RuleParser.Rule{ApplyFunc: extractRegex(Parse.CheckboxFalseTag, `- \[ ?\] (.*)`)}

	UnorderedListItem = RuleParser.Rule{ApplyFunc: extractRegex(Parse.UnorderedListTag, `- (.*)`)}
	OrderedListItem   = RuleParser.Rule{ApplyFunc: extractRegex(Parse.OrderedListTag, `\d+\. (.*)`)}
)
