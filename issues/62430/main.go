package main

import (
	"log"
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
		log.Fatal("lookup F failed with err", err)
	} else {
		log.Println("lookup F success")
	}

	f := s.(func())
	f()
}
