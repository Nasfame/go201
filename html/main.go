package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sergi/go-diff/diffmatchpatch"
	"golang.org/x/net/html"
)

var message = `
<div style='font-family:arial, "helvetica neue", helvetica, sans-serif;'>Mehr Informationen</div>
`

func main() {
	tokenizer := html.NewTokenizer(strings.NewReader(message))

	dmp := diffmatchpatch.New()

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
		original := tk.String()

		unescaped := html.UnescapeString(original)

		diffs := dmp.DiffMain(original, unescaped, false)
		diffText := dmp.DiffPrettyText(diffs)

		fmt.Println("Original:")
		fmt.Println(original)
		fmt.Println("Unescaped:")
		fmt.Println(unescaped)
		fmt.Println("Diff:")
		fmt.Println(diffText)
		fmt.Println()
	}
}
