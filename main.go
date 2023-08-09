package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"os"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
)

//go:embed template.html
var htmlTemplate string

type Output struct {
	HTML template.HTML
}

func main() {
	args := os.Args[1:]

	i := args[0]
	o := args[1]

	if _, err := os.Stat(i); os.IsNotExist(err) {
		fmt.Println("file does not exist")
		os.Exit(1)
	}

	md, err := os.ReadFile(i)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var b bytes.Buffer
	buff := bufio.NewWriter(&b)

	markdown := goldmark.New(
		goldmark.WithExtensions(
			highlighting.NewHighlighting(
				highlighting.WithStyle("monokai"),
				highlighting.WithFormatOptions(
					chromahtml.WithLineNumbers(true),
				),
			),
		),
	)

	markdown.Convert(md, buff)

	out := Output{
		HTML: template.HTML(b.String()),
	}

	tmpl, err := template.New("output").Parse(htmlTemplate)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	outFile, err := os.Create(o)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = tmpl.Execute(outFile, out)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
