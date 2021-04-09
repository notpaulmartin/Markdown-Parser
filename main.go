package main

import (
	"fmt"
	"mdParser/Compiler"
	"mdParser/Parse"
	"mdParser/Parse/RuleParser"
	"mdParser/PostParser"
	"mdParser/Rules"
	"reflect"
	"strings"
)

func main() {
	// Get commandline arguments
	//args := os.Args[1:]
	//text := strings.Join(args, " ")
	//
	//fmt.Println(text)

	//fmt.Printf("%v\n", Rules.Heading1.Apply)

	inputStr := "# [*aa **#*]( hello**)\n# **hi**"

	// Split by "\n\n"
	sections := strings.Split(inputStr, "\n\n")

	// Convert to input tree
	var inputTree []Parse.ParseTree
	for _, sectionStr := range sections {
		inputTree = append(inputTree, Parse.Section(sectionStr))
	}

	// parse lines
	parsed := RuleParser.RecursiveApply(inputTree, &Rules.All)

	var previousParsed []Parse.ParseTree

	for !reflect.DeepEqual(parsed, previousParsed) {
		// previousParsed = copy(parsed)
		previousParsed = make([]Parse.ParseTree, len(parsed))
		copy(previousParsed, parsed)

		// parse again (only using formatters, as lines will already have been parsed)
		parsed = RuleParser.RecursiveApply(parsed, &Rules.Formatters)
	}

	parsed = PostParser.Clean(parsed)
	html := Compiler.ToHtml(parsed)

	fmt.Println(parsed)
	fmt.Println(html)
}
