package main

import (
	"plugin"
	"regexp/syntax"
)

func main() {
	if syntax.Perl != 212 {
		panic("Unexpected flags")
	}

	p, err := plugin.Open("p1.so")
	if err != nil {
		panic(err)
	}
	s, err := p.Lookup("F")
	if err != nil {
		panic(err)
	}

	f := s.(func())
	f()
}
