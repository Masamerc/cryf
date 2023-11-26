package main

import (
	"fmt"

	"github.com/Masamerc/cryf/dev/dummy"
	"github.com/alexeyco/simpletable"
)

func dummyGen(generator dummy.Generator) {
	record := generator.Generate()
	record.Preview()
}

func main() {
	st := simpletable.New()
	fmt.Printf("%T\n", st)

	generator := dummy.NewSalesRecordGenerator("all")
	dummyGen(generator)

	for i := 0; i < 10; i++ {
		dummyGen(generator)
	}
}
