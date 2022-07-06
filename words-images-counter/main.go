package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func visit(n *html.Node, pwords, ppics *int) {

	if n.Type == html.TextNode {
		if n.Data != "script" && n.Data != "style" {
			*pwords += len(strings.Fields(n.Data))
		}

	} else if n.Type == html.ElementNode && n.Data == "img" {
		*ppics++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c, pwords, ppics)
	}
}

func countWordsAndImages(doc *html.Node) (int, int) {
	var words, pics int

	visit(doc, &words, &pics)

	return words, pics
}

func main() {
	url := "https://www.google.com/"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "parse failed: %s\n", err)
		os.Exit(-1)
	}

	words, pics := countWordsAndImages(doc)

	fmt.Printf("%d words and %d images\n", words, pics)
}
