package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var tokens = make(chan struct{}, 20)
var depth = 3

func main() {
	var n int
	worklist := make(chan []string)
	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)
	for ; n >= 0; n-- {
		for list := range worklist {
			for _, link := range list {
				if !seen[link] {
					n++
					seen[link] = true
					fmt.Println(n)
					go func(link string) {
						worklist <- crawl(link, depth)
					}(link)
				}
			}
		}
	}

}

func crawl(url string, depth int) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := Extract(url, depth)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

func Extract(url string, depth int) ([]string, error) {
	resp, err := http.Get(url)
	log.Printf("parsing url: %s\n", url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("failed with status code: %s", resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("error while parsing: %s", err.Error())
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil, depth)

	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node), depth int) {
	if depth <= 0 {
		return
	}
	if pre != nil {
		pre(n)
	}
	// Recursive call to process every and each node separately
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post, depth-1)
	}
	if post != nil {
		post(n)
	}
}
