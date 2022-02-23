package linkparser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

func Parse(r io.Reader) ([]Link, error){
	links := []Link{}
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)
	for _, node := range nodes {
		// fmt.Println(node)
		links = append(links, buildLink(node))
	}
	return links, nil
}

func buildLink(n *html.Node) Link {
	var ret Link
	for _, attributes := range n.Attr {
		if attributes.Key == "href" {
			ret.Href = attributes.Val
			break
		}
	}
	ret.Text = text(n)
	return ret
}

func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c)
	}
	return strings.Join(strings.Fields(ret), " ")
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c:= n.FirstChild; c!= nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}

type Link struct {
	Href string
	Text string
}