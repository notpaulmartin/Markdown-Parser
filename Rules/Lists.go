package Rules

import (
	"log"
	"mdParser/Parse"
	"mdParser/Parse/RuleParser"
	"reflect"
	"regexp"
	"strings"
)

const (
	ulRegexStr = ` *- .*\n?`
	olRegexStr = ` *\d+\. .*\n?`
)

var (
	Lists = RuleParser.Precedence{Order: []Parse.Applyable{
		&UnorderedList,
		&OrderedList,
	}}

	// TODO: nested checkboxes
	// TODO: allow...
	// 	- abdc
	//	  hello
	//    > xhei
	//    ``` ehainsf ```

	CheckboxChecked = RuleParser.Rule{
		TagName:   Parse.CheckboxTrueTag,
		ApplyFunc: extractRegex(Parse.CheckboxTrueTag, ` *- \[[x|X]\] (.*)`),
	}
	CheckboxUnchecked = RuleParser.Rule{
		TagName:   Parse.CheckboxFalseTag,
		ApplyFunc: extractRegex(Parse.CheckboxFalseTag, ` *- \[ ?\] (.*)`),
	}

	UnorderedList = RuleParser.Rule{ApplyFunc: ulFunc}
	OrderedList = RuleParser.Rule{ApplyFunc: olFunc}
)

func olFunc(input string) (bool, []Parse.ParseTree) {
	tagName := Parse.OrderedListTag
	regexStr := `((?:^` + olRegexStr + `)(?:(?:^` + olRegexStr + `)|(?:^  +` + ulRegexStr + `))+)`

	// Pre-process, identifying the unordered list
	success, blocksTree := extractRegex(tagName, regexStr)(input)
	if !success {
		return false, nil
	}

	var parsedTree []Parse.ParseTree
	for _, node := range blocksTree {
		if node.TagName == Parse.OrderedListTag {
			parsedTree = append(parsedTree, parseBlock(node.Children[0].Content, Parse.OrderedListTag)...)
		} else {
			parsedTree = append(parsedTree, node)
		}
	}

	return true, parsedTree
}

func ulFunc(input string) (bool, []Parse.ParseTree) {
	tagName := Parse.UnorderedListTag
	regexStr := `((?:^` + ulRegexStr + `)(?:(?:^` + ulRegexStr + `)|(?:^  +` + olRegexStr + `))+)`

	// Pre-process, identifying the unordered list
	success, blocksTree := extractRegex(tagName, regexStr)(input)
	if !success {
		return false, nil
	}

	var parsedTree []Parse.ParseTree
	for _, node := range blocksTree {
		if node.TagName == Parse.UnorderedListTag {
			parsedTree = append(parsedTree, parseBlock(node.Children[0].Content, Parse.UnorderedListTag)...)
		} else {
			parsedTree = append(parsedTree, node)
		}
	}

	return true, parsedTree
}

// Assumes input to be a multiline list (sub-lists allowed)
func parseBlock(blockStr string, blockType Parse.Tag) []Parse.ParseTree {
	// Split blockStr into lines
	itemLines := strings.Split(blockStr, "\n")

	// Parse lines
	var tree []Parse.ParseTree // Output tree

	type nestingsElem struct {
		IndentLen int
		ListType  Parse.Tag
	}
	var nestings []nestingsElem
	nestings = append(nestings, nestingsElem{
		IndentLen: 0,
		ListType:  blockType,
	})

	for _, line := range itemLines {
		ok, newListType, newIndentLen, listItem := parseListItem(line)
		if !ok {
			// Should never happen
			log.Panic("[Should never happen] Could not parse line which has been identified as a list item:", line)
		}

		prevNesting := peek(nestings).(nestingsElem)

		// Backtrace to correct indentation
		for newIndentLen < prevNesting.IndentLen {
			// Pop from nesting-stack while indent of new item < indent of previous item
			t, ts := pop(nestings)
			prevNesting = t.(nestingsElem)
			nestings = ts.([]nestingsElem)
		}

		// Is there a new indentation?
		if newIndentLen > prevNesting.IndentLen {
			// Push new indent
			nestings = append(nestings, nestingsElem{
				IndentLen: newIndentLen,
				ListType:  newListType,
			})

			prevNesting = peek(nestings).(nestingsElem)
		}

		// Is there a change in list type?
		if newListType != prevNesting.ListType {
			// Pop old list type
			_, ts := pop(nestings)
			nestings = ts.([]nestingsElem)

			// Push new list type
			nestings = append(nestings, nestingsElem{
				IndentLen: newIndentLen,
				ListType:  newListType,
			})
		}

		// Add list item wrapped in list blocks
		node := listItem
		for i := len(nestings)-1; i >= 0; i-- {
			node = Parse.ParseTree{
				TagName:  nestings[i].ListType,
				Children: Parse.UnitTree(node),
			}
		}

		tree = append(tree, node)
	}

	return tree
}

func parseListItem(line string) (success bool, listType Parse.Tag, indentLen int, listItem Parse.ParseTree) {
	ulRegex := regexp.MustCompile(`^( *)- (.*)$`)     // Unordered List item
	olRegex := regexp.MustCompile(`^( *)\d+\. (.*)$`) // Ordered List item

	if ulRegex.MatchString(line) { // if line is an unordered list item
		// Parse item
		matches := ulRegex.FindStringSubmatch(line)
		indentLen := len(matches[1]) // Match group 1 (indentation)
		itemContent := matches[2]    // Match group 2 (list item text)

		return true, Parse.UnorderedListTag, indentLen, Parse.ParseTree{
			TagName:  Parse.ListItemTag,
			Children: Parse.RawChild(itemContent),
		}
	} else if olRegex.MatchString(line) { // if line is an unordered list item
		// Parse item
		matches := olRegex.FindStringSubmatch(line)
		indentLen := len(matches[1]) // Match group 1 (indentation)
		itemContent := matches[2]    // Match group 2 (list item text)

		return true, Parse.OrderedListTag, indentLen, Parse.ParseTree{
			TagName:  Parse.ListItemTag,
			Children: Parse.RawChild(itemContent),
		}
	}

	return false, "", 0, Parse.ParseTree{}
}

/// Stack actions ///
//  pop(stack []T) (T, []T)
func pop(stack interface{}) (interface{}, interface{}) {
	if reflect.TypeOf(stack).Kind() != reflect.Slice {
		log.Panic("Pop can only handle slices. Non-slice type used: ", reflect.ValueOf(stack).Kind())
	}

	s := reflect.ValueOf(stack)
	newStack := s.Slice(0, s.Len()-1).Interface()
	return peek(stack), newStack
}

//  peek(stack []T) (T)
func peek(stack interface{}) interface{} {
	if reflect.TypeOf(stack).Kind() != reflect.Slice {
		log.Panic("Peek can only handle slices. Non-slice type used: ", reflect.ValueOf(stack).Kind())
	}

	s := reflect.ValueOf(stack)
	return s.Index(s.Len() - 1).Interface()
}
