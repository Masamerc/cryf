package main

import (
	"fmt"

	"github.com/Masamerc/cryf/dev/dummy"
	"github.com/alexeyco/simpletable"
)

func dummyGen(generator dummy.Generator) {
	record := generator.Generate()
	record2 := generator.Generate()
	record.Preview()
	record2.Preview()
}

func main() {
	st := simpletable.New()
	fmt.Printf("%T\n", st)

	generator := dummy.NewSalesRecordGenerator("all")
	dummyGen(generator)

}
