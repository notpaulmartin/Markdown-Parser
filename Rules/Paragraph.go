package Rules

import (
	"mdParser/Parse"
	"mdParser/Parse/RuleParser"
	"regexp"
	"strings"
)

var (
	Paragraph = RuleParser.Rule{
		TagName:   Parse.ParagraphTag,
		ApplyFunc: applyParagraph,
	}
)

func applyParagraph(input string) (bool, []Parse.ParseTree) {
	// Executed after every other line("extractor")-rule, so matches any line that hasn't been matched otherwise
	// Splits paragraphs by empty line after (\n\n) and captures the last paragraph without needing "\n\n" after it
	regexStr := `^(.+)\n\n|^(.+)\z` // "\z" = Absolute end of string

	r, err := regexp.Compile(singleline + ungreedy + regexStr)
	if err != nil || !r.MatchString(input) {
		return false, nil
	}

	// Compile matches into nodes
	matches := r.FindAllStringSubmatch(input, -1)
	var matchNodes []Parse.ParseTree
	for _, match := range matches {
		paragraphText := selectNonEmpty(match[1:])  // match = (string, capture-groups...)
		matchNodes = append(matchNodes, Parse.ParseTree{
			TagName:  Parse.ParagraphTag,
			Children: Parse.RawChild(paragraphText),
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
	//return false, nil
}
