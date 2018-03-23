package main

import (
	"github.com/GUMI-golang/gumi"
)

func main() {
	data, err := gumi.Request("file:///./ex/example.gumi.xml")
	if err != nil {
		panic(err)
	}
	par := gumi.NewParser(data)
	par.Verbose = true
	for par.Step() == nil{

	}

}