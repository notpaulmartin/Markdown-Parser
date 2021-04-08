package Rules

import "mdParser/Parse"

type applyFunc func(string) (bool, []Parse.ParseTree)

type Rule struct {
	ApplyFunc applyFunc
}

// Wrapper around ApplyFunc, to allow for interface matching
func (r *Rule) Apply(x string) (bool, []Parse.ParseTree) {
	return (*r).ApplyFunc(x)
}
