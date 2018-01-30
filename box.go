package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/atotto/clipboard"
)

func init() {
	// Suppress standard logging information
	log.SetFlags(0)
}

func main() {
	if len(os.Args) == 1 {
		log.Fatal("no input provided")
	}

	input := strings.Join(os.Args[1:], " ")
	box := generate(input)

	fmt.Println(box)

	if err := clipboard.WriteAll(box); err != nil {
		log.Fatalf("error copying to clipboard: %v", err)
	}
}

func generate(input string) string {
	topBottom := strings.Repeat("*", len(input)+3)

	buf := new(bytes.Buffer)
	buf.WriteString("/")
	buf.WriteString(topBottom)
	buf.WriteString("\n")
	buf.WriteString("* ")
	buf.WriteString(input)
	buf.WriteString(" *")
	buf.WriteString("\n")
	buf.WriteString(topBottom)
	buf.WriteString("/")

	return buf.String()
}
