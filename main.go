package main

import (
	"fmt"
	"github.com/terwer/api-middleware-go/api/endpoint/markdown"
	"os"
)

func main() {
	var md = "# Hello, World!"
	content, err := os.ReadFile("testdata/test.md")
	//content, err := os.ReadFile("testdata/test.html")
	if err != nil {
		fmt.Println("Unable to read file")
		return
	}
	md = string(content)

	//params := markdown.RequestBody{
	//	To: markdown.HTML,
	//	Md: md,
	//}

	//params := markdown.RequestBody{
	//	To: markdown.MD,
	//	Md: md,
	//}

	params := markdown.RequestBody{
		To: markdown.DOM,
		Md: md,
	}

	//params := markdown.RequestBody{
	//	To: markdown.TEXT,
	//	Md: md,
	//}

	result, err := markdown.HandleMarkdown(&params)
	if err != nil {
		fmt.Printf("Rendering result encountered an error: %s\n", err.Error())
	} else {
		fmt.Printf("Render result is: %s\n", result.Data)
	}
}
