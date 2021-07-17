package Rules

import (
	"github.com/notpaulmartin/mdParser/Parse"
	"github.com/notpaulmartin/mdParser/Parse/RuleParser"
)

var (
	// TODO: fix
	EscapeCharacters = RuleParser.Precedence{Order: []Parse.Applyable{
		&EscMinus,
		&EscHash,
		&EscStar,
		&EscUnderscore,
	}}

	EscMinus      = RuleParser.Rule{ApplyFunc: extractRegex(Parse.TextTag, `\\(-)`)}
	EscHash       = RuleParser.Rule{ApplyFunc: extractRegex(Parse.TextTag, `\\(#)`)}
	EscStar       = RuleParser.Rule{ApplyFunc: extractRegex(Parse.TextTag, `\\(*)`)}
	EscUnderscore = RuleParser.Rule{ApplyFunc: extractRegex(Parse.TextTag, `\\(_)`)}
)
