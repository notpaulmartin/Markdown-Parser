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

	CheckboxChecked = RuleParser.Rule{
		TagName: Parse.CheckboxTrueTag,
		ApplyFunc: extractRegex(Parse.CheckboxTrueTag, `- \[[x|X]\] (.*)`),
	}
	CheckboxUnchecked = RuleParser.Rule{
		TagName: Parse.CheckboxFalseTag,
		ApplyFunc: extractRegex(Parse.CheckboxFalseTag, `- \[ ?\] (.*)`),
	}

	UnorderedListItem = RuleParser.Rule{
		TagName: Parse.UnorderedListTag,
		ApplyFunc: extractRegex(Parse.UnorderedListTag, `- (.*)`),
	}
	OrderedListItem   = RuleParser.Rule{
		TagName: Parse.OrderedListTag,
		ApplyFunc: extractRegex(Parse.OrderedListTag, `\d+\. (.*)`),
	}
)
