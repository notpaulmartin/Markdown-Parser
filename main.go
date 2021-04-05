package main

import (
	"fmt"
	"mdParser/Rules"
)

func main() {
	// Get commandline arguments
	//args := os.Args[1:]
	//text := strings.Join(args, " ")
	//
	//fmt.Println(text)

	success, h1 := Rules.UnorderedListItem.Apply("- aa ***hello*** **xy**z")
	fmt.Println("success:", success)
	fmt.Println("format:", h1)
}
