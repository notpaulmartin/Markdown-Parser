package main

import (
	"fmt"
	"mdParser/Parse"
	"mdParser/Rules"
	"reflect"
)

func main() {
	// Get commandline arguments
	//args := os.Args[1:]
	//text := strings.Join(args, " ")
	//
	//fmt.Println(text)

	input := "[aa **#]( hello**)"

	// TODO: split by...
	//  - "\n\n"
	//  -

	// parse lines
	_, parsed := Rules.All.Apply(input)
	var previousParsed []Parse.ParseTree

	for !reflect.DeepEqual(parsed, previousParsed) {
		// previousParsed = copy(parsed)
		previousParsed = make([]Parse.ParseTree, len(parsed))
		copy(previousParsed, parsed)

		// parse again (only using formatters, as lines will already have been parsed)
		parsed = Rules.RecursiveApply(parsed, &Rules.Formatters)
	}

	// TODO (??): implement post-processor
	//  - to join codeblocks that have been split by intermediate "\n\n"

	// TODO: implement compiler to HTML

	fmt.Println(parsed)
}
