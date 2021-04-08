package Rules

import "mdParser/Parse/RuleParser"

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

/// Text - If everything else fails
var Text = RuleParser.Rule{ApplyFunc: func(s string) (bool, []Parse.ParseTree) {
	return true, []Parse.ParseTree{Parse.Text(s)}
}}
