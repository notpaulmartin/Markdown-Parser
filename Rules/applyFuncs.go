// Helper functions that handle common patterns for parsing md
package Rules

import (
	"mdParser/Parse"
	"regexp"
)

// E.g. "# abc"  ->  "abc"
func extractRegex(tagName Parse.Tag, regex string) applyFunc {
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
func extractRegexWContent(tagName Parse.Tag, regex string, content string) applyFunc {
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
func applyRegexInText(tagName Parse.Tag, regex string) applyFunc {
	return func(input string) (bool, []Parse.ParseTree) {
		// (?U) is the "Ungreedy" RegEx-modifier
		r, err := regexp.Compile(`(?U)` + regex)
		if err != nil || !r.MatchString(input) {
			return false, nil
		}

		matches := r.FindStringSubmatch(input)[1:]

		// tagsNo := len(matches)*2+1
		tags := make([]Parse.ParseTree, 0)

		// Interlace raw strings with tags:
		// 	RAW [match] RAW [match] RAW
		raws := r.Split(input, -1)
		for i, rawStr := range raws {
			if rawStr == "" {
				continue
			}
			tags = append(tags, Parse.Raw(rawStr))

			if i >= len(matches) {
				continue
			}
			tags = append(tags, Parse.ParseTree{
				TagName:  tagName,
				Children: Parse.RawChild(matches[i]),
			})
		}

		return true, tags
	}
}
