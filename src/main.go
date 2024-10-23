package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	outputFile := flag.String("o", "", "Output HTML file (default: same as input with .html extension)")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Usage: go run main.go <markdown_file> [-o <output_file>]")
		return
	}

	markdownFile := flag.Args()[0]
	htmlContent, err := ConvertMarkdownToHTML(markdownFile)
	if err != nil {
		fmt.Printf("Error converting markdown to HTML: %v\n", err)
		return
	}

	var output string
	if *outputFile != "" {
		output = *outputFile
	} else {
		output = changeExtension(markdownFile, ".html")
	}

	err = os.WriteFile(output, []byte(htmlContent), 0644)
	if err != nil {
		fmt.Printf("Error writing HTML to file: %v\n", err)
		return
	}

	fmt.Printf("HTML written to %s\n", output)
}

func changeExtension(filename, newExt string) string {
	return filename[:len(filename)-len(filepath.Ext(filename))] + newExt
}
