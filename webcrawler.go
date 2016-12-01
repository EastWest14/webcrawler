package main

import (
	"fmt"
	"golang.org/x/net/html"
	//"io/ioutil"
	"net/http"
)

const (
	ADDR = "http://google.com"
)

func main() {
	resp, err := http.Get(ADDR)
	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err.Error())
	//}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	for _, link := range visit(doc) {
		fmt.Println(link)
	}
}

func visit(node *html.Node) (links []string) {
	links = []string{}

	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attribute := range node.Attr {
			if attribute.Key == "href" {
				links = append(links, attribute.Val)
			}
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, visit(c)...)
	}

	return links
}
