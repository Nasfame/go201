package main

import (
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"os"
	"strings"
)

var message = `
<div style='font-family:arial, "helvetica neue", helvetica, sans-serif;'>Mehr Informationen</div>
`

func main() {
	tokenizer := html.NewTokenizer(strings.NewReader(message))

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

		unescaped := html.UnescapeString(tk.String())
		fmt.Println(unescaped)
		
	}

}
