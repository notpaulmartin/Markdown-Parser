package Rules

import (
	"mdParser/Parse"
	"regexp"
)

func applyRegex(tagName string, regex string) applyFunc {
	return func(input string) (bool, []Parse.Tag) {
		r, err := regexp.Compile(regex)
		if err != nil || !r.MatchString(input) {return false, nil}

		matches := r.FindStringSubmatch(input)
		match := matches[len(matches)-1]


		return true, []Parse.Tag{{
			Name:    tagName,
			Content: match,
		}}
	}
}

func applyRegex2(tagName string, regex string) applyFunc {
	return func (input string) (bool, []Parse.Tag) {
		// (?U) is the "Ungreedy" RegEx-modifier
		r, err := regexp.Compile(`(?U)` + regex)
		if err != nil || !r.MatchString(input) {
			return false, nil
		}

		matches := r.FindStringSubmatch(input)[1:]

		// tagsNo := len(matches)*2+1
		tags := make([]Parse.Tag, 0)

		// Interlace raw strings with tags:
		// 	RAW [match] RAW [match] RAW
		raws := r.Split(input, -1)
		for i, rawStr := range raws {
			if rawStr == "" {continue}
			tags = append(tags, Parse.Raw(rawStr))

			if i >= len(matches) {continue}
			tags = append(tags, Parse.Tag{tagName, matches[i]})
		}

		return true, tags
	}
}

