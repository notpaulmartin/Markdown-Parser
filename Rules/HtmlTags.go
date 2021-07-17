package Rules

import (
	"github.com/notpaulmartin/mdParser/Parse"
	"github.com/notpaulmartin/mdParser/Parse/RuleParser"
	"regexp"
	"strings"
)

var (
	HtmlTag = RuleParser.Rule{
		TagName:   Parse.HtmlTagTag,
		ApplyFunc: applyHtmlTag,
	}
)

func applyHtmlTag(input string) (bool, []Parse.ParseTree) {
	// TODO
	regexStr := `<(.*)>(.*)</.*>`

	r, err := regexp.Compile(singleline + ungreedy + regexStr)
	if err != nil || !r.MatchString(input) {
		return false, nil
	}

	// Compile matches into nodes
	matches := r.FindAllStringSubmatch(input, -1)
	var matchNodes []Parse.ParseTree
	for _, match := range matches {
		// match = (string, capture-groups...)
		htmlTagName := match[1]
		htmlTagContent := match[2]
		matchNodes = append(matchNodes, Parse.ParseTree{
			TagName:  Parse.HtmlTagTag,
			Children: Parse.RawChild(htmlTagContent),
			Content: htmlTagName,
		})
	}

	// Compile non-matches into RAW-nodes
	raws := r.Split(input, -1)
	var rawNodes []Parse.ParseTree
	for _, rawStr := range raws {
		rawNodes = append(rawNodes, Parse.Raw(rawStr))
	}

	if len(raws) <= 0 {
		return true, matchNodes
	}

	var tree []Parse.ParseTree
	if strings.HasPrefix(input, raws[0]) {
		tree = interlace(rawNodes, matchNodes)
	} else {
		tree = interlace(matchNodes, rawNodes)
	}

	return true, tree
}
