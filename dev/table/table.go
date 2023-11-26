package table

// import (
// 	"github.com/alexeyco/simpletable"
// )

// // using simpletable take all maps from a slice and display them in a table

// func ShowTable(items []map[string]string) {
// 	table := simpletable.New()

// 	// use the keys of the first map as header
// 	header := make([]string, 0, len(items[0]))
// 	for k := range items[0] {
// 		header = append(header, k)
// 	}

// 	// set header
// 	table.Header = &simpletable.Header{
// 		Cells: make([]*simpletable.Cell, 0, len(header)),
// 	}

// 	for _, v := range header {
// 		table.Header.Cells = append(table.Header.Cells, &simpletable.Cell{
// 			Align: simpletable.AlignCenter,
// 			Text:  v,
// 		})
// 	}

// 	// populate table
// 	for _, item := range items {
// 		r := make([]*simpletable.Cell, 0, len(item))

// 		for _, key := range header {
// 			r = append(r, &simpletable.Cell{
// 				Align: simpletable.AlignCenter,
// 				Text:  item[key],
// 			})
// 		}
// 		table.Body.Cells = append(table.Body.Cells, r)
// 	}

// 	table.Print()

// }
