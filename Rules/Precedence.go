package Rules

import "mdParser/Parse"

type Precedence struct {
	order []Applyable
}

func (p *Precedence) Apply(input string) (bool, []Parse.ParseTree)  {
	for _, rule := range p.order {
		success, parsed := rule.Apply(input)
		if success {return success, parsed}
	}

	return false, nil
}
