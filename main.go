package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"mdParser/Compiler"
	"mdParser/Parse"
	"mdParser/Parse/RuleParser"
	"mdParser/PostParser"
	"mdParser/PrettyPrinter"
	"mdParser/Rules"
	"os"
	"reflect"
)

func main() {
	// Get filename from commandline arguments
	args := os.Args[1:]
	fname := args[0]

	// Read input Markdown string from file
	inputBytes, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Panic("Could not read input file")
	}
	inputStr := string(inputBytes)

	// Split Markdown string by "\n\n"
	//sections := strings.Split(inputStr, "\n\n")

	// Convert sections to input tree
	var inputTree []Parse.ParseTree
	//for _, sectionStr := range sections {
	//	inputTree = append(inputTree, Parse.Section(sectionStr))
	//}

	inputTree = append(inputTree, Parse.Raw(inputStr))

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

	parsed = AddParagraphs(parsed)

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
	PrettyPrinter.PrettyPrint(parsed)

	// Convert to HTML
	html := Compiler.ToHtml(parsed)
	fmt.Println(html)
}
