package Rules

import "mdParser/Parse"

type Applyable interface {
	Apply(string) (bool, []Parse.ParseTree)
}
