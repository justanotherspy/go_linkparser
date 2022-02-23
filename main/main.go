package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	filename := flag.String("file", "ex1.html", "the html file to parse for links")
	flag.Parse()
	f, err := os.Open(*filename)
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	defer f.Close()
}
