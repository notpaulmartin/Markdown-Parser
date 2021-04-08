package main

import (
	"fmt"
	"mdParser/Parse"
	"mdParser/Parse/RuleParser"
	"mdParser/PostParser"
	"mdParser/Rules"
	"reflect"
)

func main() {
	// Get commandline arguments
	//args := os.Args[1:]
	//text := strings.Join(args, " ")
	//
	//fmt.Println(text)

	//fmt.Printf("%v\n", Rules.Heading1.Apply)

	input := "# [*aa **#*]( hello**)"
	//input := "*a* *x*"

	// TODO: split by...
	//  - "\n\n"
	//  -

	// parse lines
	_, parsed := Rules.All.Apply(input)
	var previousParsed []Parse.ParseTree

	for !reflect.DeepEqual(parsed, previousParsed) {
		fmt.Println(parsed)

		// previousParsed = copy(parsed)
		previousParsed = make([]Parse.ParseTree, len(parsed))
		copy(previousParsed, parsed)

		// parse again (only using formatters, as lines will already have been parsed)
		parsed = RuleParser.RecursiveApply(parsed, &Rules.Formatters)
	}

	parsed = PostParser.Clean(parsed)

	// TODO (??): implement post-processor
	//  - to join codeblocks that have been split by intermediate "\n\n"

	// TODO: implement compiler to HTML

	fmt.Printf("%v\n", parsed)
}
