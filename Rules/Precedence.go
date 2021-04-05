package Rules

import "mdParser/Parse"

type Precedence struct {
	order []Rule
}

func (p *Precedence) Apply(input string) (bool, []Parse.Tag)  {
	for _, rule := range p.order {
		success, parsed := rule.Apply(input)
		if success {return success, parsed}
	}

	return false, nil
}
