// Helper functions that handle common patterns for parsing md
package Rules

import (
	"github.com/notpaulmartin/mdParser/Parse"
	"github.com/notpaulmartin/mdParser/Parse/RuleParser"
	"regexp"
	"strings"
)

// Regex flags
const (
	ungreedy  = `(?Um)`
	multiline = `(?m)`
	singleline = `(?s)`
)


// Meant for entire lines (Greedy regex)
// E.g. "# abc"  ->  "abc"
func extractRegex(tagName Parse.Tag, regex string) RuleParser.ApplyFunc {
	return func(input string) (bool, []Parse.ParseTree) {
		r, err := regexp.Compile(multiline + regex)
		if err != nil || !r.MatchString(input) {
			return false, nil
		}

		// Compile matches into nodes
		matches := r.FindAllStringSubmatch(input, -1)
		matchNodes := matchesToNodes(tagName, matches, "")

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

// Meant for patterns within a line (Ungreedy regex)
// E.g. "a*b*c"  ->  "a", Italics{"b"}, "c"
func applyRegexInText(tagName Parse.Tag, regex string) RuleParser.ApplyFunc {
	return extractRegex(tagName, ungreedy + regex)
}

func matchesToNodes(tagName Parse.Tag, matches [][]string, content string) []Parse.ParseTree {
	// Compile matches into nodes
	var matchNodes []Parse.ParseTree
	for _, match := range matches {
		// Select first non-empty capture group, as regex or-statements (a)|(b) may result in multiple groups
		matchStr := selectNonEmpty(match[1:]) // match = (string, capture-groups)
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
			// out.append( xs.popTag() )
			x, xs = xs[0], xs[1:]
			out = append(out, x)
		}
		if len(ys) > 0 {
			// out.append( ys.popTag() )
			y, ys = ys[0], ys[1:]
			out = append(out, y)
		}
	}
}

// Select the first non-empty string from a list
func selectNonEmpty(strs []string) string {
	for _, str := range strs {
		if str != "" {
			return str
		}
	}

	return ""
}
