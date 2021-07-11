package Rules

import (
	"mdParser/Parse"
	"mdParser/Parse/RuleParser"
)

var (
	HtmlTags = RuleParser.Rule{
		TagName:   Parse.HtmlTagTag,
		ApplyFunc: applyHtmlTag,
	}
)

func applyHtmlTag(input string) (bool, []Parse.ParseTree) {
	//regexStr := `<(.*)>(.*)</(.*)>(.*)`
	return false, nil
}
