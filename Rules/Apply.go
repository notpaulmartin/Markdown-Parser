package Rules

import (
	"mdParser/Parse"
	"regexp"
)

func RecursiveApply(tree []Parse.ParseTree, applyable Applyable) []Parse.ParseTree {
	// Base case: empty tree
	if len(tree) == 0 {
		return tree
	}

	// Base case: unit tree and reached unexpanded node
	if len(tree) == 1 && tree[0].TagName == Parse.RawTag {
		success, parse := applyable.Apply(tree[0].Content) // Try to expand node
		if !success {
			// If can't expand, convert to Text leaf node
			return Parse.UnitTree(Parse.Text(tree[0].Content))
		}
		return parse
	}

	// Base case: unit tree and reached leaf
	if len(tree) == 1 && len(tree[0].Children) == 0 {
		return tree
	}

	// Unit tree and at expandable node  ->  Go deeper
	if len(tree) == 1 {
		tree[0].Children = RecursiveApply(tree[0].Children, applyable)
		return tree
	}

	// Not a unit tree  ->  Recurse over all root nodes
	var newTree []Parse.ParseTree
	for _, node := range tree {
		newTree = append(newTree, RecursiveApply(Parse.UnitTree(node), applyable)...)
	}
	return newTree
}

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
