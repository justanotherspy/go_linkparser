package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	linkparser "dansdomain.net/html_link_parser"
)

func main() {
	filename := flag.String("file", "ex3.html", "the html file to parse for links")
	flag.Parse()
	f, err := os.Open(*filename)
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	defer f.Close()
	
	links, err := linkparser.Parse(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", links)
}
