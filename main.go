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

	input := "# aa **# hello**"

	_, parsed := Rules.All.Apply(input)
	var previousParsed []Parse.ParseTree

	for !reflect.DeepEqual(parsed, previousParsed) {
		// previousParsed = copy(parsed)
		previousParsed = make([]Parse.ParseTree, len(parsed))
		copy(previousParsed, parsed)

		// parse again
		parsed = Rules.RecursiveApply(parsed, &Rules.Formatters)
	}

	fmt.Println(parsed)
}
