package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

var (
	words map[string]struct{}
)

func main() {
	words := make(map[string]struct{}, 0)
	twl, err := Asset("static/twl2014")
	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.NewBuffer(twl)
	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		w := scanner.Text()
		words[w] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	scanner = bufio.NewScanner(os.Stdin)

	fmt.Println("Enter words")
	for scanner.Scan() {
		w := scanner.Text()

		if _, exists := words[w]; exists {
			fmt.Println(w, "is a word")
		} else {
			fmt.Println(w, "is triflute")
		}
	}
}
