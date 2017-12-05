package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("no input provided")
	}

	input := strings.Join(os.Args[1:], " ")
	tb := strings.Repeat("*", len(input)+3)

	buf := new(bytes.Buffer)
	buf.WriteString("/")
	buf.WriteString(tb)
	buf.WriteString("\n")
	buf.WriteString("* ")
	buf.WriteString(input)
	buf.WriteString(" *")
	buf.WriteString("\n")
	buf.WriteString(tb)
	buf.WriteString("/")

	fmt.Println(buf.String())
}
