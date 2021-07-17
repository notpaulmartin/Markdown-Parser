# Markdown Parser

⚠️ **WARNING:** Do not use in a production environment. For a better Go Markdown parser, consider [goldmark](https://github.com/yuin/goldmark).

This parser is a personal project, built out of curiosity and fun. I created it to compile my CV from Markdown to Html (and then to Pdf with some styling).

## Usage example
```go
package main

import (
    "fmt"
    "github.com/notpaulmartin/mdParser"
)

func main() {
    markdown := "# Hello\nWorld"
    html := mdParser.MdToHtml(markdown)
    fmt.Println(html)
}
```
[Open example on goplay.space](https://goplay.space/#cAPmvUR7K5V)
