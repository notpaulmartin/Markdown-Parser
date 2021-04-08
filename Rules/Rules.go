package Rules

import (
	"mdParser/Parse"
	"mdParser/Parse/RuleParser"
)

/* TODO:
 *  - Multi-line rules
 *  	-> Code-blocks (tab / ```)
 *  -
*/

var All = RuleParser.Precedence{Order: []Parse.Applyable{
	&Extractors,
	&Formatters,
	&Text,
}}

var Extractors = RuleParser.Precedence{Order: []Parse.Applyable{
	&Headings,
	&Lists,
}}

var Formatters = RuleParser.Precedence{Order: []Parse.Applyable{
	&Links,
	&TextFormatting,
	&Text,
}}

/// Text - If everything else fails
var Text = RuleParser.Rule{ApplyFunc: func(s string) (bool, []Parse.ParseTree) {
	return true, []Parse.ParseTree{Parse.Text(s)}
}}
