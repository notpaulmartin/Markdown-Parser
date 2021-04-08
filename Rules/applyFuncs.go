// Helper functions that handle common patterns for parsing md
package Rules

import (
	"mdParser/Parse"
	"mdParser/Parse/RuleParser"
	"regexp"
	"strings"
)

// Regex flags
const (
	ungreedy  = `(?Um)`
	multiline = `(?m)`
)


// E.g. "# abc"  ->  "abc"
func extractRegex(tagName Parse.Tag, regex string) RuleParser.ApplyFunc {
	return func(input string) (bool, []Parse.ParseTree) {
		r, err := regexp.Compile(multiline + regex)
		if err != nil || !r.MatchString(input) {
			return false, nil
		}

		matches := r.FindAllStringSubmatch(input, -1)
		matchNodes := matchesToNodes(tagName, matches, "")  // Compile matches into nodes
		return true, matchNodes
	}
}

// E.g. "- [ ] abc"  ->  Checkbox{ "abc", false }
func extractRegexWContent(tagName Parse.Tag, regex string, content string) RuleParser.ApplyFunc {
	return func(input string) (bool, []Parse.ParseTree) {
		r, err := regexp.Compile(multiline + regex)
		if err != nil || !r.MatchString(input) {
			return false, nil
		}

		matches := r.FindAllStringSubmatch(input, -1)
		matchNodes := matchesToNodes(tagName, matches, content)  // Compile matches into nodes
		return true, matchNodes
	}
}

// E.g. "a*b*c"  ->  "a", Italics{"b"}, "c"
func applyRegexInText(tagName Parse.Tag, regex string) RuleParser.ApplyFunc {
	return func(input string) (bool, []Parse.ParseTree) {
		r, err := regexp.Compile(ungreedy + regex)
		if err != nil || !r.MatchString(input) {
			return false, nil
		}

		matches := r.FindAllStringSubmatch(input, -1)
		raws := r.Split(input, -1)

		// Compile matches into nodes
		matchNodes := matchesToNodes(tagName, matches, "")

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

func matchesToNodes(tagName Parse.Tag, matches [][]string, content string) []Parse.ParseTree {
	// Compile matches into nodes
	var matchNodes []Parse.ParseTree
	for _, match := range matches {
		matchStr := match[1] // match = (string, capture-group)
		matchNodes = append(matchNodes, Parse.ParseTree{
			TagName:  tagName,
			Children: Parse.RawChild(matchStr),
			Content: content,
		})
	}

	return matchNodes
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
