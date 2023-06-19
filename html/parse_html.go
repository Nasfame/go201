package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var message = `
<div style='font-family: "Comic Sans MS", Times, serif;'>Mehr Informationen</div>
`

func main() {
	tokenizer := html.NewTokenizer(strings.NewReader(message))

	fmt.Println("original message", message)

	for {
		tokenizer.Next()

		if err := tokenizer.Err(); err != nil {
			if errors.Is(err, io.EOF) {
				// end of the file, break out of the loop
				break
			} else {
				os.Exit(1)
			}
		}

		tk := tokenizer.Token()
		fmt.Println(tk.String())

		//fmt.Println(html.UnescapeString(tk.String()))

	}
}
