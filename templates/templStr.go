package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/template"
)

func main() {
	tmplStr := `Hello, {{.Name}}! You are {{.Age}} years old.`
	tmpl := template.New("x")

	tmpl, err := tmpl.Parse(tmplStr)

	name := "\bard"
	fmt.Printf("%v\n",name)
	jsonName, err := json.Marshal(name)

	fmt.Println(string(jsonName),err)

	// Create some data to pass to the template.
	data := map[string]string{
		"Name": name,
		"Age":  "1",
	}

	// Execute the template with the data.
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println(err)
		return
	}
}
