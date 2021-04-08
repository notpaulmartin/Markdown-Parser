// Helper functions that handle common patterns for parsing md
package Rules

import (
	"fmt"
	"mdParser/Parse"
	"mdParser/Parse/RuleParser"
	"regexp"
	"strings"
)


// E.g. "# abc"  ->  "abc"
func extractRegex(tagName Parse.Tag, regex string) RuleParser.ApplyFunc {
	return func(input string) (bool, []Parse.ParseTree) {
		r, err := regexp.Compile(regex)
		if err != nil || !r.MatchString(input) {
			return false, nil
		}

		matches := r.FindStringSubmatch(input)
		match := matches[len(matches)-1]

		return true, []Parse.ParseTree{{
			TagName:  tagName,
			Children: Parse.RawChild(match),
		}}
	}
}

// E.g. "- [ ] abc"  ->  Checkbox{ "abc", false }
func extractRegexWContent(tagName Parse.Tag, regex string, content string) RuleParser.ApplyFunc {
	return func(input string) (bool, []Parse.ParseTree) {
		r, err := regexp.Compile(regex)
		if err != nil || !r.MatchString(input) {
			return false, nil
		}

		matches := r.FindStringSubmatch(input)
		match := matches[len(matches)-1]

		return true, []Parse.ParseTree{{
			TagName:  tagName,
			Children: Parse.RawChild(match),
			Content:  content,
		}}
	}
}

// E.g. "a*b*c"  ->  "a", Italics{"b"}, "c"
func applyRegexInText(tagName Parse.Tag, regex string) RuleParser.ApplyFunc {
	return func(input string) (bool, []Parse.ParseTree) {
		// (?U) is the "Ungreedy" RegEx-modifier
		r, err := regexp.Compile(`(?U)` + regex)
		if err != nil || !r.MatchString(input) {
			return false, nil
		}

		matches := r.FindAllStringSubmatch(input, -1)
		raws := r.Split(input, -1)

		// Compile matches into nodes
		var matchNodes []Parse.ParseTree
		for _, match := range matches {
			matchStr := match[1] // match = (string, capture-group)
			matchNodes = append(matchNodes, Parse.ParseTree{
				TagName:  tagName,
				Children: Parse.RawChild(matchStr),
			})
		}

		fmt.Println("->", matches)

		// Compile non-matches into RAW-nodes
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
}

// Interlace two parseTrees: [x, y, x, y, …]
func interlace(xs, ys []Parse.ParseTree) []Parse.ParseTree {
	var out []Parse.ParseTree
	var x, y Parse.ParseTree
	for {
		if len(xs)+len(ys) <= 0 {
			return out
		}

		if len(xs) > 0 {
			// out.append( xs.pop() )
			x, xs = xs[0], xs[1:]
			out = append(out, x)
		}
		if len(ys) > 0 {
			// out.append( ys.pop() )
			y, ys = ys[0], ys[1:]
			out = append(out, y)
		}
	}
}
