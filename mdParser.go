package mdParser

import (
	"mdParser/Compiler"
	"mdParser/Parse"
	"mdParser/Parse/RuleParser"
	"mdParser/PostParser"
	"mdParser/Rules"
	"reflect"
)

func MdToHtml(md string) (html string) {
	var inputTree []Parse.ParseTree
	inputTree = Parse.RawChild(md)

	// First only parse entire lines (until can't parse anymore)
	var parsed, previousParsed []Parse.ParseTree
	parsed = inputTree
	for !reflect.DeepEqual(parsed, previousParsed) {
		// previousParsed = copy(parsed)
		previousParsed = make([]Parse.ParseTree, len(parsed))
		copy(previousParsed, parsed)

		//parsed = PostParser.Trim(parsed)
		// parse again (only using formatters, as lines will already have been parsed)
		parsed = RuleParser.RecursiveApply(parsed, &Rules.Extractors)
	}

	// All lines that aren't anything else are paragraphs
	parsed = addParagraphs(parsed)

	// Then parse only inline formatting (until can't parse anymore)
	previousParsed = nil
	for !reflect.DeepEqual(parsed, previousParsed) {
		// previousParsed = copy(parsed)
		previousParsed = make([]Parse.ParseTree, len(parsed))
		copy(previousParsed, parsed)

		// parsed = PostParser.Trim(parsed)
		// parse again (only using formatters, as lines will already have been parsed)
		parsed = RuleParser.RecursiveApply(parsed, &Rules.Formatters)
	}

	// Clean up tree ("Fix it in post")
	parsed = PostParser.Clean(parsed)
	// PrettyPrinter.PrettyPrint(parsed)

	// Convert to HTML
	html = Compiler.ToHtml(parsed)
	return html
}
