package dummy

type SalesRecord struct {
	Date     string
	Id       string
	Amount   int
	Location *Location
	Item     *Item
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
