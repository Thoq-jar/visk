package main

import (
	"github.com/russross/blackfriday/v2"
	"os"
)

func ConvertMarkdownToHTML(markdownFile string) (string, error) {
	data, err := os.ReadFile(markdownFile)
	if err != nil {
		return "", err
	}

	html := blackfriday.Run(data)
	return string(html), nil
}
