package dummy

import "fmt"

type SalesRecord struct {
	Date     string
	Id       string
	Amount   int
	Location *Location
	Item     *Item
}

func (r SalesRecord) Preview() {
	fmt.Println("Date: ", r.Date)
	fmt.Println("Id: ", r.Id)
	fmt.Println("Amount: ", r.Amount)
	fmt.Println("Location:")
	fmt.Println("  City: ", r.Location.City)
	fmt.Println("  State: ", r.Location.State)
	fmt.Println("  Storename: ", r.Location.StoreName)
	fmt.Println("Item: ")
	fmt.Println("  Category: ", r.Item.Category)
	fmt.Println()
}

type Location struct {
	City      string
	State     string
	StoreName string
}

type Item struct {
	Category string
}

type SalesRecordGenerator struct {
	dateRange string
}

func NewSalesRecordGenerator(dateRange string) *SalesRecordGenerator {
	return &SalesRecordGenerator{
		dateRange: dateRange,
	}
}

func (g SalesRecordGenerator) Generate() Record {
	return SalesRecord{
		Date:   "dummy",
		Id:     "dummy",
		Amount: 100,
		Location: &Location{
			City:      "dummy",
			State:     "dummy",
			StoreName: "dummy",
		},
		Item: &Item{
			Category: "food",
		},
	}

}
