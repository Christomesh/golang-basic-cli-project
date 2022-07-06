package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func getAllWebsiteLinks(links []string, n *html.Node) []string {

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = getAllWebsiteLinks(links, c)
	}
	return links
}

func main() {
	url := "https://www.google.com/"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Errorf("getting %s: %s", url, resp.StatusCode)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	result := getAllWebsiteLinks(nil, doc)

	fmt.Println(result)

}
