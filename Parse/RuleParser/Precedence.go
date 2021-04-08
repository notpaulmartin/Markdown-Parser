package RuleParser

import "mdParser/Parse"

// Defines an Order (precedence) in which rules and precedences should be applied
type Precedence struct {
	Order []Parse.Applyable
}

func (p *Precedence) Apply(input string) (bool, []Parse.ParseTree)  {
	for _, rule := range p.Order {
		success, parsed := rule.Apply(input)
		if success {return success, parsed}
	}

	return false, nil
}
