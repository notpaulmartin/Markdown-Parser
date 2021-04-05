package Rules

import (
	"mdParser/Parse"
)

type applyFunc func(string) (bool, []Parse.Tag)

type Rule struct {
	Apply applyFunc
}

//func newRegexRule(tag string, regex string) Rule {
//
//	applyRegex :=
//		func(input string) (bool, []Parse.Tag) {
//			r, err := regexp.Compile(regex)
//
//			// Error
//			if err != nil {return false, nil}
//			// No match
//			if !r.MatchString(input) {return false, nil}
//
//			captured := r.FindStringSubmatch(input)[1:]
//
//
//
//			return true, nil
//		}
//	return Rule{applyRegex}
//}
