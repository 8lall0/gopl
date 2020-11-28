package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		var attr string

		switch n.Data {
		case "a", "link":
			attr = "href"
		case "script", "img":
			attr = "src"
		default:
			attr = ""
		}

		if attr != "" {
			for _, a := range n.Attr {
				if a.Key == attr {
					links = append(links, a.Val)
				}
			}
		}
	}

	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}

	return links
}
