package RuleParser

import "mdParser/Parse"

type ApplyFunc func(string) (bool, []Parse.ParseTree)

type Rule struct {
	TagName   Parse.Tag
	ApplyFunc ApplyFunc
}

// Wrapper around ApplyFunc, to allow for interface matching
func (r *Rule) Apply(x string) (bool, []Parse.ParseTree) {
	return (*r).ApplyFunc(x)
}
