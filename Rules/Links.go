package Rules

import (
	"github.com/notpaulmartin/mdParser/Parse"
	"github.com/notpaulmartin/mdParser/Parse/RuleParser"
	"regexp"
	"strings"
)

// Links and Images

var (
	Links = RuleParser.Precedence{Order: []Parse.Applyable{
		&Img,
		&Link,
	}}

	Img  = RuleParser.Rule{TagName: Parse.ImgTag, ApplyFunc: linkApply(Parse.ImgTag, `\!(\[.*\]\(.*\))`)}
	Link = RuleParser.Rule{TagName: Parse.LinkTag, ApplyFunc: linkApply(Parse.LinkTag, `(\[.*\]\(.*\))`)}
)

func linkApply(tagName Parse.Tag, regex string) RuleParser.ApplyFunc {
	return func(input string) (bool, []Parse.ParseTree) {
		// (?U) is the "Ungreedy" RegEx-modifier
		r, err := regexp.Compile(ungreedy + regex) // Match all links
		if err != nil || !r.MatchString(input) {
			return false, nil
		}

		matches := r.FindAllStringSubmatch(input, -1)
		raws := r.Split(input, -1)

		// Compile matches into nodes
		var linkNodes []Parse.ParseTree
		for _, match := range matches {
			matchStr := match[1] // match = (string, capture-group)
			success, text, linkUrl := parseLink(matchStr)
			if !success {
				return false, nil
			}

			linkNodes = append(linkNodes, Parse.ParseTree{
				TagName:  tagName,
				Children: Parse.RawChild(text),
				Content:  linkUrl,
			})
		}

		// Compile non-matches into RAW-nodes
		var rawNodes []Parse.ParseTree
		for _, rawStr := range raws {
			rawNodes = append(rawNodes, Parse.Raw(rawStr))
		}

		if len(raws) <= 0 {
			return true, linkNodes
		}

		var tree []Parse.ParseTree
		if strings.HasPrefix(input, raws[0]) {
			tree = interlace(rawNodes, linkNodes)
		} else {
			tree = interlace(linkNodes, rawNodes)
		}

		return true, tree
	}
}

// Extract text and url from md link
func parseLink(link string) (success bool, text, url string) {
	r := regexp.MustCompile(`(?U)^\[(.*)\]\((.*)\)$`)

	matches := r.FindStringSubmatch(link)[1:]
	if len(matches) <= 1 {
		return false, "", ""
	}

	return true, matches[0], matches[1]
}
