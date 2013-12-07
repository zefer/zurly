package main

import (
	"github.com/subosito/iglo"
	"log"
	"flag"
	"os"
)

var outFile = flag.String("out", "index.html", "Filename of the HTML output")
var inFile = flag.String("in", "api.md", "Filename of the API blueprint input")

func main() {
	flag.Parse()

	f, err := os.Open(*inFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w, err := os.Create(*outFile)
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	err = iglo.MarkdownToHTML(w, f)
	if err != nil {
		log.Fatal(err)
	}
}
