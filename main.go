package mdParser

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

	html := MdToHtml(inputStr)
	fmt.Println(html)
}
