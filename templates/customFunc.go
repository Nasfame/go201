package main

import (
	"bytes"
	"fmt"
	"text/template"
)

func greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func main() {
	templateStr := `
Greetings! {{ greet .name }}
`

	t := template.New("my_template")
	t.Funcs(template.FuncMap{
		"greet": greet,
	})

	// Parse the template.
	tmpl, err := t.Parse(templateStr)
	if err != nil {
		panic(err)
	}

	var templateBuf bytes.Buffer

	// Execute the template.
	err = tmpl.Execute(&templateBuf, map[string]string{
		"name": "hiro",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(templateBuf.String())
}
